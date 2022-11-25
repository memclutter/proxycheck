package proxycheck

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestProxyRequest(t *testing.T) {
	proxyURLEnv := os.Getenv("PROXY_URL")
	targetURLEnv := os.Getenv("TARGET_URL")
	if len(proxyURLEnv) == 0 {
		t.Skip("Skipped because PROXY_URL not set")
		return
	} else if len(targetURLEnv) == 0 {
		t.Skip("Skipped because TARGET_URL not set")
		return
	}

	timeout := 10 * time.Second
	proxyURL, err := url.Parse(proxyURLEnv)
	require.NoError(t, err, "must be correct PROXY_URL url address")

	body, err := ProxyRequest(targetURLEnv, proxyURL, timeout)
	require.NoError(t, err, "must be success proxy request")
	assert.Contains(t, string(body), "Welcome to nginx!", "must be contain Welcome to nginx!")
}

func TestProxyRequestNot200StatusCode(t *testing.T) {
	proxyURLEnv := os.Getenv("PROXY_URL")
	targetURLEnv := os.Getenv("TARGET_URL")
	if len(proxyURLEnv) == 0 {
		t.Skip("Skipped because PROXY_URL not set")
		return
	} else if len(targetURLEnv) == 0 {
		t.Skip("Skipped because TARGET_URL not set")
		return
	}

	timeout := 10 * time.Second
	proxyURL, err := url.Parse(proxyURLEnv)
	require.NoError(t, err, "must be correct PROXY_URL url address")

	_, err = ProxyRequest(targetURLEnv+"/not_found_page/", proxyURL, timeout)
	require.ErrorContains(t, err, `http error, status code 404`, "must be contain error")
}

func TestProxyRequestNotReachableError(t *testing.T) {
	timeout := 100 * time.Millisecond
	proxyURL, err := url.Parse("http://localhost:1111")
	require.NoError(t, err, "must be correct PROXY_URL url address")

	_, err = ProxyRequest("http://localhost", proxyURL, timeout)
	require.ErrorContains(t, err, `can't http get: Get "http://localhost":`, "must be error")
}
