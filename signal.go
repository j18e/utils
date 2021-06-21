package utils

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

// ListenForSignal captures os.Signals and returns an error when one of them is
// received.
func ListenForSignal(ctx context.Context, sigs ...os.Signal) error {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, sigs...)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case sig := <-sigchan:
		return fmt.Errorf("received signal %s", sig)
	}
}
