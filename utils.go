package proxycheck

import (
	"context"
	"fmt"
	"h12.io/socks"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func createProxyTransport(proxyURL *url.URL, timeout time.Duration) *http.Transport {
	httpTransport := &http.Transport{
		TLSHandshakeTimeout:   timeout,
		IdleConnTimeout:       timeout,
		ResponseHeaderTimeout: timeout,
		DisableKeepAlives:     true,
		DisableCompression:    false,
	}
	if strings.HasPrefix(proxyURL.Scheme, "socks") {
		httpTransport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return socks.Dial(fmt.Sprintf("%s?timeout=%s", proxyURL.String(), timeout.String()))(network, addr)
		}
	} else {
		httpTransport.Proxy = http.ProxyURL(proxyURL)
		httpTransport.DialContext = (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: timeout,
		}).DialContext
	}
	return httpTransport
}

func readResponse(httpResponse *http.Response) ([]byte, error) {
	defer httpResponse.Body.Close()
	return ioutil.ReadAll(httpResponse.Body)
}
