package proxycheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestApp_UnknownJudge is the regression test for the --judge flag: an unknown
// judge name must abort with an error before any proxy is checked. A deliberately
// unparsable address keeps the run off the network if the guard is missing.
func TestApp_UnknownJudge(t *testing.T) {
	err := NewApp().Run([]string{"proxycheck", "--judge", "foo", "not-an-addr"})
	assert.Error(t, err, "unknown judge must cause a non-nil error")
	if err != nil {
		assert.Contains(t, err.Error(), "unknown judge: foo")
	}
}
