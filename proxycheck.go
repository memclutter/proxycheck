package proxycheck

import (
	"fmt"
	"h12.io/socks"
	"net/http"
	"net/url"
)

type checkerOpts struct {
	TargetURL string
	ProxyURL  *url.URL
}

type checker func(opts *checkerOpts) error

var checkerMap = map[string]checker{
	"http":   checkHttp,
	"https":  checkHttp,
	"socks4": checkSocks,
	"socks5": checkSocks,
}

func check(targetURL, proxyAddr string) error {
	proxyURL, err := url.Parse(proxyAddr)
	if err != nil {
		return fmt.Errorf("invalid proxy address %s", proxyAddr)
	}

	if checkerFunc, ok := checkerMap[proxyURL.Scheme]; ok {
		return checkerFunc(&checkerOpts{TargetURL: targetURL, ProxyURL: proxyURL})
	} else {
		return fmt.Errorf("unknown proxy scheme %s", proxyURL.Scheme)
	}
}

func checkSocks(opts *checkerOpts) error {
	dial := socks.Dial(opts.ProxyURL.String())
	transport := &http.Transport{Dial: dial}
	client := &http.Client{Transport: transport}

	resp, err := client.Get(opts.TargetURL)
	if err != nil {
		return err
	} else if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("status %s", resp.Status)
	} else {
		return nil
	}
}

func checkHttp(opts *checkerOpts) error {
	transport := &http.Transport{Proxy: http.ProxyURL(opts.ProxyURL)}
	client := &http.Client{Transport: transport}

	resp, err := client.Get(opts.TargetURL)
	if err != nil {
		return err
	} else if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("status %s", resp.Status)
	} else {
		return nil
	}
}
