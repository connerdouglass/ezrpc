package ezrpc

import (
	"context"
	"net/http"
)

// New creates a new, empty, root middleware on which to build
func New() Middleware {
	return Middleware(func(ctx context.Context, r *http.Request) (context.Context, error) {
		return ctx, nil
	})
}

// Middleware is a piece of middleware functionality in the RPC pipeline that is
// executed before a handler is executed.
type Middleware func(ctx context.Context, r *http.Request) (context.Context, error)

func thenOne(f Middleware, next Middleware) Middleware {
	return Middleware(func(ctx context.Context, r *http.Request) (context.Context, error) {
		ctx, err := f(ctx, r)
		if err != nil {
			return ctx, err
		}
		return next(ctx, r)
	})
}

// Then appends another middleware to the pipeline and returns a new pipeline tail
func (f Middleware) Then(next ...Middleware) Middleware {
	m := Middleware(f)
	for _, n := range next {
		m = thenOne(m, n)
	}
	return m
}

// Handle adds a HTTP handler as a sink to the middleware pipeline
func (f Middleware) Handle(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		// Apply the middleware
		ctx := r.Context()
		ctx, err := f(ctx, r)
		if err != nil {
			writeError(rw, err)
			return
		}

		// Update the context on the request
		r = r.WithContext(ctx)

		// Then pass to the handler function
		handler.ServeHTTP(rw, r)

	})
}
