package proxycheck

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestProxyAddrToIpPort(t *testing.T) {
	testingTable := []struct {
		title        string
		proxyAddr    string
		exceptedIP   net.IP
		exceptedPort int64
	}{
		{
			title:        "must be correct convert ip:port format",
			proxyAddr:    "89.23.43.100:8899",
			exceptedIP:   net.ParseIP("89.23.43.100"),
			exceptedPort: 8899,
		},
		{
			title:        "must be correct convert ip without port format",
			proxyAddr:    "86.100.202.1",
			exceptedIP:   net.ParseIP("86.100.202.1"),
			exceptedPort: 80,
		},
	}

	for _, table := range testingTable {
		t.Run(table.title, func(t *testing.T) {
			ip, port, err := proxyAddrToIpPort(table.proxyAddr)

			require.NoError(t, err, "must be run without errors")
			assert.Equal(t, table.exceptedIP, ip, "must be equal ip")
			assert.Equal(t, table.exceptedPort, port, "must be equal port")
		})
	}
}

func TestReadResponse(t *testing.T) {
	exceptedBody := []byte(`test`)
	resp := &http.Response{Body: io.NopCloser(bytes.NewReader(exceptedBody))}
	actualBody, err := readResponse(resp)

	assert.NoError(t, err, "must be ok")
	assert.Equal(t, exceptedBody, actualBody, "must be body equal")
}

func TestCreateProxyTransport(t *testing.T) {
	testingTable := []struct {
		title    string
		proxyURL *url.URL
		timeout  time.Duration
	}{
		{
			title: "socks4 proxy use correct http transport and timeout",
			proxyURL: &url.URL{
				Scheme: "socks4",
				Host:   "100.200.100.20:500",
			},
			timeout: 4 * time.Second,
		},
		{
			title: "http proxy use correct http transport and timeout",
			proxyURL: &url.URL{
				Scheme: "http",
				Host:   "239.10.200.30",
			},
			timeout: 120 * time.Millisecond,
		},
	}

	for _, table := range testingTable {
		t.Run(table.title, func(t *testing.T) {
			httpTransport := createProxyTransport(table.proxyURL, table.timeout)

			assert.Equal(t, table.timeout, httpTransport.TLSHandshakeTimeout, "must be equal timeout")
			assert.Equal(t, table.timeout, httpTransport.IdleConnTimeout, "must be equal timeout")
			assert.Equal(t, table.timeout, httpTransport.ResponseHeaderTimeout, "must be equal timeout")
			assert.True(t, httpTransport.DisableKeepAlives, "must be disable keep alive")
			assert.False(t, httpTransport.DisableCompression, "must be enable compression")

			assert.NotNil(t, httpTransport.DialContext, "must be set dialer")

			// Only for non socks proxy, check proxy url correct
			if table.proxyURL.Scheme == "http" || table.proxyURL.Scheme == "https" {
				exceptedProxyURL, _ := http.ProxyURL(table.proxyURL)(nil)
				actualProxyURL, _ := httpTransport.Proxy(nil)

				assert.Equal(t, exceptedProxyURL, actualProxyURL)
			}
		})
	}
}
