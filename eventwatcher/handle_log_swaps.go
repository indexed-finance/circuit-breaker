package eventwatcher

import (
	"context"
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/indexed-finance/circuit-breaker/alerts"
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
	ec *ethclient.Client,
	poolAddress common.Address,
) error {
	for {
		select {
		case <-ctx.Done():
			wc.swapSub.Unsubscribe()
			close(wc.swapCh)
			close(wc.notifSwapCh)
			return nil
		case err := <-wc.swapSub.Err():
			wc.logger.Error("event subscription error received", zap.Error(err))
			wc.swapSub.Unsubscribe()
			close(wc.swapCh)
			close(wc.notifSwapCh)
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
					)

					// lock the authorizer since bind.TransactOpts is not threadsafe
					authorizer.Lock()

					gasPrice, err := utils.GetGasPrice(ctx, wc.backend, wc.minimumGwei, wc.gasMultiplier)
					if err != nil {
						wc.logger.Error("failed to suggest gasprice", zap.Error(err))
					} else {
						wc.logger.Info("gas price calculated (includes boost)", zap.String("gas.price", gasPrice.String()))
						wc.setPublicSwap(ctx, authorizer, poolAddress, gasPrice, wc.name, ec)
					}

					// we need to unset the gas price that we overrode the transactor with
					// so that future uses of this transactor have the gas price set to nil
					authorizer.TransactOpts.GasPrice = nil
					authorizer.Unlock()

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