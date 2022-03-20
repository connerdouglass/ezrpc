package main

import (
	"context"
	"log"
	"net/http"

	"github.com/connerdouglass/ezrpc"
)

func main() {

	// Create a pipeline of middleware
	rpc := ezrpc.
		New().
		Then(
			ezrpc.Middleware(func(ctx context.Context, _ *http.Request) (context.Context, error) {
				log.Println("First middleware")
				return ctx, nil
			}),
			ezrpc.Middleware(func(ctx context.Context, _ *http.Request) (context.Context, error) {
				log.Println("Second middleware")
				return ctx, nil
			}),
		)

	// Create an HTTP mux and register the RPC hooks
	mux := http.NewServeMux()
	mux.Handle("/v1/sum", rpc.Handle(ezrpc.Handle(Sum)))

	// Serve the HTTP mux
	http.ListenAndServe("127.0.0.1:8080", mux)

}
