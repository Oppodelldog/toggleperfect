package remote

import (
	"fmt"
	"time"
)

const waitTimeoutMs = 200

func printTimeoutMessage(asset string) {
	fmt.Printf("timeout while syncing %s - connect remote client or set TP_ENABLE_REMOTE_UI=false\n", asset)
}

func newOutputTimeout() *time.Timer {
	return time.NewTimer(time.Millisecond * waitTimeoutMs)
}
