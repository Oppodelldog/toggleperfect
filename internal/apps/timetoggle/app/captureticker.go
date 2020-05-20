package app

import (
	"context"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/app/repo"
	"log"
	"time"
)

func NewCaptureTicker(ctx context.Context, captureID string) CaptureTicker {
	tickerContext, cancelFunc := context.WithCancel(ctx)

	return CaptureTicker{
		ctx:       tickerContext,
		cancel:    cancelFunc,
		captureID: captureID,
	}
}

type CaptureTicker struct {
	ctx       context.Context
	cancel    context.CancelFunc
	captureID string
}

func (ct CaptureTicker) Start() {
	go func() {
		t := time.NewTicker(time.Second * 30)
		for {
			select {
			case <-t.C:
				latestStop := repo.Capture{
					ID:        ct.captureID,
					Timestamp: time.Now().Unix(),
				}
				err := repo.SetLatestStop(latestStop)
				if err != nil {
					log.Printf("error settings latest stop time on '%s': %v", ct.captureID, err)
				}
			case <-ct.ctx.Done():
				log.Printf("CaptureTicker stopped for '%s'", ct.captureID)
				return
			}
		}
	}()
}

func (ct CaptureTicker) Stop() {
	ct.cancel()
}
