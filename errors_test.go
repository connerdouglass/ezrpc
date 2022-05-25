package ezrpc_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/connerdouglass/ezrpc"
	"github.com/stretchr/testify/assert"
)

func TestErrorWithCode(t *testing.T) {
	err := errors.New("socket closed")
	err1 := ezrpc.ErrorWithCode(err, http.StatusBadRequest, "saving to database")
	assert.Equal(t, "saving to database: socket closed", err1.Error())
}
