package ezrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Handler is the type for a function that handles an RPC request and returns a response
type Handler[T, K any] func(context.Context, *T) (*K, error)

// Handle wraps an RPC handler in HTTP logic so that it can be connected to an HTTP server
func Handle[T, K any](handler Handler[T, K]) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		// Decode the request from JSON
		req, err := decodeRequestBody[T](r)
		if err != nil {
			writeError(rw, ErrorWithCode(err, http.StatusBadRequest))
			return
		}

		// Get the context from the request
		ctx := r.Context()

		// Call the handler function
		res, err := handler(ctx, req)
		if err != nil {
			writeError(rw, err)
			return
		}

		// Send the response
		writeResponse(rw, http.StatusOK, res)

	})
}

// decodeRequestBody decodes the incoming request body into the appropriate type for an RPC handler.
func decodeRequestBody[T any](r *http.Request) (*T, error) {
	if r.Body == nil {
		r.Body = io.NopCloser(bytes.NewBufferString(""))
	}
	// TODO: allow other types of decoders than just JSON
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var req T
	if err := dec.Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

// writeResponse writes the given response data and status code.
func writeResponse(rw http.ResponseWriter, statusCode int, data any) {

	// Convert the response data to JSON. Fallback to a marshal error if that doesn't work
	// for some reason.
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		statusCode = http.StatusInternalServerError
		jsonBytes, _ = json.Marshal(map[string]any{
			"error": fmt.Sprintf("response marshal error: %s", err.Error()),
		})
	}

	// Write the header and the data
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	rw.Write(jsonBytes)

}

// writeError is a convenience function that writes an error response
func writeError(rw http.ResponseWriter, err error) {

	// Determine the status code
	statusCode := http.StatusBadRequest
	if err, ok := err.(errorWithCode); ok {
		statusCode = err.Code()
	}

	// Write the error with the appropriate code
	writeResponse(rw, statusCode, map[string]any{
		"error": err.Error(),
	})

}
