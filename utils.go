package proxycheck

import (
	"context"
	"fmt"
	"h12.io/socks"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func proxyAddrToIpPort(proxyAddr string) (ip net.IP, port int64, err error) {
	parts := strings.Split(proxyAddr, ":")
	if len(parts) == 2 {
		port, err = strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return ip, port, fmt.Errorf("invalid port: %v", err)
		}
	} else if len(parts) == 1 {
		port = 80
	}

	ip = net.ParseIP(parts[0])
	if len(ip) == 0 {
		return ip, port, fmt.Errorf("invalid ip '%s'", parts[0])
	}

	return
}

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
