package utils_test

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/j18e/utils"
)

func ExampleHTTPServer() {
	// use a logger of your choice (Logrus is a good option)
	var l *logger

	// add an HTTP server with a health check
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	srv := utils.HTTPServer{Addr: ":3000", Handler: mux, Logger: l}

	// create a context that times out after one second
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// run the http server whose context will time out
	fmt.Println(srv.ListenAndServe(ctx, time.Second))
	// Output: context deadline exceeded
}
