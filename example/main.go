package main

import (
	"net/http"

	"github.com/connerdouglass/ezrpc"
)

func main() {

	// Create an HTTP mux and register the RPC hooks
	mux := http.NewServeMux()
	mux.Handle("/v1/sum", ezrpc.Handle(Sum))

	// Serve the HTTP mux
	http.ListenAndServe("127.0.0.1:8080", mux)

}
