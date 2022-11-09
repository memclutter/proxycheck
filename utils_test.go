package proxycheck

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestReadResponse(t *testing.T) {
	exceptedBody := []byte(`test`)
	resp := &http.Response{Body: io.NopCloser(bytes.NewReader(exceptedBody))}
	actualBody, err := readResponse(resp)

	assert.NoError(t, err, "must be ok")
	assert.Equal(t, exceptedBody, actualBody, "must be body equal")
}
