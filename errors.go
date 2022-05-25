package ezrpc

import "fmt"

// errorWithCode defines the internal interface of an error that also contains an error code.
// It is useful to wrap errors in this interface within your RPC hooks to specify the type of
// error that occurred, so the appropriate HTTP status code can be returned.
type errorWithCode interface {
	error
	Code() int
}

// ErrorWithCode wraps an error with an HTTP status code, so that it will be used as the RPC response.
func ErrorWithCode(err error, code int, msgAndArgs ...any) error {
	if len(msgAndArgs) > 0 {
		msg := fmt.Sprint(msgAndArgs...)
		err = fmt.Errorf("%s: %s", msg, err)
	}
	return &_implErrorWithCode{
		err,
		code,
	}
}

// implErrorWithCode is the internal implementation of `errorWithCode`. By convention, the underscore
// speficies that this type should be file-private.
type _implErrorWithCode struct {
	err  error
	code int
}

func (err *_implErrorWithCode) Code() int {
	return err.code
}

func (err *_implErrorWithCode) Error() string {
	return err.err.Error()
}
