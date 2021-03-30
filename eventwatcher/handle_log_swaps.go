package eventwatcher

import (
	"context"
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexed-finance/circuit-breaker/alerts"
	"github.com/indexed-finance/circuit-breaker/circuit"
	"github.com/indexed-finance/circuit-breaker/database"
	"github.com/indexed-finance/circuit-breaker/utils"
	"go.uber.org/zap"
)

func (wc *WatchedContract) handleLogSwaps(
	ctx context.Context,
	db *database.Database,
	alerter *alerts.Alerter,
	authorizer *utils.Authorizer,
	breakPercentage float64,
	poolAddress common.Address,
) error {
	defer func() {
		wc.swapSub.Unsubscribe()
		close(wc.swapCh)
		close(wc.notifSwapCh)
	}()
	for {
		select {
		case <-ctx.Done():
			wc.logger.Info("context cancelled, exiting")
			return nil
		case err := <-wc.swapSub.Err():
			wc.logger.Error("event subscription error received", zap.Error(err))
			return err
		case evLog := <-wc.swapCh:
			wc.logger.Debug("received new event (log swap)", zap.Any("event", evLog))
			// skip reverted logs
			if evLog.Raw.Removed {
				wc.logger.Warn("event was marked as removed by blockchain due to reorg")
				continue
			}

			// send event to watchers
			select {
			case wc.notifSwapCh <- evLog:
			default:
			}

			prevSwapFee, err := wc.getPreviousSwapFee(evLog.Raw.BlockNumber - 1)
			if err != nil {
				wc.logger.Error("failed to get previous swap fee", zap.Error(err))
				continue
			}

			currSpotPrice, err := wc.getCurrentSpotPrice(evLog)
			if err != nil {
				wc.logger.Error("failed to get current spot price", zap.Error(err))
				continue

			}
			wc.logger.Debug(
				"LOG_SWAP event spot price calculated",
				zap.String("tx.hash", evLog.Raw.TxHash.String()),
				zap.String("price", currSpotPrice.String()),
				zap.String("token.out", evLog.TokenOut.String()),
				zap.String("token_amount.out", evLog.TokenAmountOut.String()),
				zap.String("token.in", evLog.TokenIn.String()),
				zap.String("token_amount.in", evLog.TokenAmountIn.String()),
				zap.Uint64("block.number", evLog.Raw.BlockNumber),
			)

			prevSpotPrice, err := wc.getPreviousSpotPrice(db, evLog, prevSwapFee)
			if err != nil {
				wc.logger.Error("failed to calculate previous block spot price", zap.Error(err))
				continue
			}

			// multipliy by 100 to get correct change percentage
			change := (utils.GetChangePercentageBig(currSpotPrice, prevSpotPrice)) * 100
			if change != 0 {
				wc.logger.Warn(
					"spot price fluctuation detected",
					zap.String("price.previous", prevSpotPrice.String()),
					zap.String("price.current", currSpotPrice.String()),
					zap.String("token.out", evLog.TokenOut.String()),
					zap.String("token.in", evLog.TokenIn.String()),
					zap.String("tx.hash", evLog.Raw.TxHash.String()),
					zap.Uint64("block_number.current", evLog.Raw.BlockNumber),
					zap.Float64("price.fluctuation", change),
				)
				// only handle breaking if the price decreases (change < 0)
				// and if the change is outside the break precentage window
				// to do this we calculate the absolute value of the change percentage
				if math.Abs(change) >= breakPercentage {

					// if the circuit is already broken there is no need to send another transaction
					wc.brokenLock.Lock()
					if wc.broken {
						wc.logger.Warn(
							"price fluctuation outside of acceptable bounds, but circuit already broken this indicates we had some delayed events",
						)
						wc.brokenLock.Unlock()
						continue
					}
					wc.brokenLock.Unlock()

					wc.logger.Warn(
						"price fluctuation outside of acceptable bounds, breaking!",
						zap.Float64("change.abs", math.Abs(change)),
						zap.Float64("break.percentage", breakPercentage),
					)

					// check to see if an anctive circuit break is ongoing
					wc.breakLock.Lock()
					if wc.breaking {
						wc.logger.Warn("attempting to trigger circuit break while one is ongoing")
						wc.breakLock.Unlock()
						continue
					}
					wc.breaking = true
					wc.breakLock.Unlock()
					if err := circuit.BreakCircuit(circuit.BreakConfig{
						Ctx:           ctx,
						BC:            wc.bc,
						MinimumGwei:   wc.minimumGwei,
						GasMultiplier: wc.gasMultiplier,
						Auth:          authorizer,
						PoolAddress:   poolAddress,
						PoolName:      wc.name,
						Logger:        wc.logger,
						Breaker:       wc.breaker,
					}); err != nil {
						wc.logger.Error(
							"circuit break failed to execute properly",
							zap.Error(err),
							zap.String("pool", wc.name),
						)
					}
					// unset the breaking boolean
					wc.breakLock.Lock()
					wc.breaking = false
					wc.breakLock.Unlock()

					if err := alerter.NotifyCircuitBreak(
						fmt.Sprintf(
							"a circuit break is in progress. pool: %s, tokenIn: %s, tokenOut: %s, changePercent: %0.2f, reason: spot price fluctuation",
							wc.name, evLog.TokenIn.String(), evLog.TokenOut.String(), change,
						),
					); err != nil {
						wc.logger.Error("circuit break notification failed")
					}
				}
			}

		}
	}
}
