package utils_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/j18e/utils"
)

func ExampleRunner() {
	// use a logger of your choice (Logrus is a good option)
	var l *logger

	// create a new Runner
	var r utils.Runner

	// make the runner listen for ctrl+c
	r.AddSignalListener(os.Interrupt)

	// add an HTTP server with a health check
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	r.Add(func(ctx context.Context) error {
		srv := utils.HTTPServer{Addr: ":3000", Handler: mux, Logger: l}
		return srv.ListenAndServe(ctx, time.Second)
	})

	// add a custom function that
	r.Add(func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.NewTimer(time.Millisecond).C:
			return errors.New("an error occurred")
		}
	})

	fmt.Println(r.Run(l))
	// Output: an error occurred
}
