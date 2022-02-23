# ezrpc

`ezrpc` is a super small Go library that uses the new Go 1.18 generics to enable RPC functionality.

I built this for myself so I can separate my application logic from gateway/HTTP logic.

## Example

Easily connect RPCs to your existing HTTP router:

```go
func main() {

	// Create an HTTP mux and register the RPC hooks
	mux := http.NewServeMux()
	mux.Handle("/v1/sum", ezrpc.Handle(Sum))

	// Serve the HTTP mux
	http.ListenAndServe("127.0.0.1:8080", mux)

}
```

Define request and response types, and a handler function for each RPC hook:

```go
type sumRequest struct {
	Numbers []int `json:"numbers"`
}

type sumResponse struct {
	Sum int `json:"sum"`
}

// Sum calculates the sum of a slice of numbers
func Sum(ctx context.Context, req *sumRequest) (*sumResponse, error) {
	var sum int
	for _, number := range req.Numbers {
		sum += number
	}
	return &sumResponse{Sum: sum}, nil
}
```

## Contributions

Contributions are welcome and encouraged. Please feel free to submit PRs!

Some features that would be useful but don't yet exist:

- Middleware
- Optional support for non-JSON encodings (XML, YAML, etc.)
- Storing more metadata in `context` value
