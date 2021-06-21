package utils

import (
	"context"
	"net/http"
	"time"
)

// HTTPServer is a wrapper around http.Server which shuts down the server in
// the event of a cancelled context.
type HTTPServer struct {
	Addr    string
	Handler http.Handler
	Logger  Logger
}

// ListenAndServe starts the http.Server running and listens for the given
// context being cancelled. If cancelled, a server shutdown is initiated with
// the given timeout.
func (s HTTPServer) ListenAndServe(ctx context.Context, shutdownTimeout time.Duration) error {
	srv := http.Server{Addr: s.Addr, Handler: s.Handler}
	errchan := make(chan error, 1)
	go func() { errchan <- srv.ListenAndServe() }()

	select {
	case err := <-errchan:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		s.Logger.Info("shutting down http server")
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return ctx.Err()
	}
}
