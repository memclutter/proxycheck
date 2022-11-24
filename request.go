package proxycheck

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ProxyRequest
// Performs an HTTP request through the proxy specified in the proxyUrl
func ProxyRequest(target string, proxyURL *url.URL, timeout time.Duration) ([]byte, error) {
	body := make([]byte, 0)
	httpClient := http.Client{
		Timeout:   timeout,
		Transport: createProxyTransport(proxyURL, timeout),
	}

	if resp, err := httpClient.Get(target); err != nil {
		return body, fmt.Errorf("can't http get: %v", err)
	} else if body, err := readResponse(resp); err != nil {
		return body, fmt.Errorf("error read response: %v", err)
	} else if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("http error, status code %d, response %s", resp.StatusCode, body)
	} else {
		return body, nil
	}
}
