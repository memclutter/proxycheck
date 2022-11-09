package proxycheck

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestProxyRequest(t *testing.T) {
	proxyURLEnv := os.Getenv("PROXY_URL")
	if len(proxyURLEnv) == 0 {
		t.Skip("Skipped because PROXY_URL not set")
		return
	}

	timeout := 10 * time.Second
	proxyURL, err := url.Parse(proxyURLEnv)
	require.NoError(t, err, "must be correct PROXY_URL url address")

	body, err := ProxyRequest("http://www.wfuchs.de/azenv.php", proxyURL, timeout)
	require.NoError(t, err, "must be success proxy request")
	assert.Contains(t, fmt.Sprintf("%s", body), "REQUEST_URI = /azenv.php", "must be contain REQUEST_URI")
}
