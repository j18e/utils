package utils

import (
	"context"
	"os"
)

// Runner is a group of functions which should be executed concurrently and
// monitored for return errors.
type Runner struct {
	funcs []func(context.Context) error
}

// Add adds a new function to the Runner.
func (r *Runner) Add(fn func(context.Context) error) {
	r.funcs = append(r.funcs, fn)
}

// AddSignalListener is a shortcut for adding ListenForSignal to the runner's
// jobs.
func (r *Runner) AddSignalListener(sigs ...os.Signal) {
	r.Add(func(ctx context.Context) error {
		return ListenForSignal(ctx, sigs...)
	})
}

// Run starts each of the Runner's functions and listens for errors on each of
// them. Once the first error comes in the context is cancelled and Run waits
// for all of the functions to return.
func (r *Runner) Run(logger Logger) error {
	errchan := make(chan error, len(r.funcs))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for _, fn := range r.funcs {
		fn := fn
		go func() {
			errchan <- fn(ctx)
		}()
	}
	err := <-errchan
	if logger != nil {
		logger.Errorf("shutting down due to: %s", err)
	}
	cancel()
	for i := 0; i < len(r.funcs)-1; i++ {
		<-errchan
	}
	return err
}
