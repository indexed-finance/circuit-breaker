package eventwatcher

import (
	"context"

	"go.uber.org/zap"
)

func (wc *WatchedContract) handleSwapToggles(ctx context.Context) {
	defer func() {
		wc.toggleSub.Unsubscribe()
		close(wc.notifToggleCh)
		close(wc.toggleCh)
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case err := <-wc.toggleSub.Err():
			wc.logger.Error("toggle subscription error received", zap.Error(err))
			return
		case evLog := <-wc.toggleCh:
			wc.logger.Debug("received new event (swap toggle)", zap.Any("event", evLog))
			// skip reverted logs
			if evLog.Raw.Removed {
				wc.logger.Warn("event was marked as removed by blockchain due to reorg")
				continue
			}

			select {
			case wc.notifToggleCh <- evLog:
			default:
			}

			wc.brokenLock.Lock()
			if !evLog.Enabled { // indicates swap has been disabled
				wc.broken = true
			} else {
				wc.broken = false
			}
			wc.brokenLock.Unlock()
		}
	}
}
