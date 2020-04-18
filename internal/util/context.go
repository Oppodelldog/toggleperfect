package util

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func NewInterruptContext() context.Context {

	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		for {
			select {
			case <-c:
				cancelFunc()
				return
			}
		}
	}()

	return ctx
}
