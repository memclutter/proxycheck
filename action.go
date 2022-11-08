package proxycheck

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/urfave/cli/v2"
	"h12.io/socks"
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

func Action(c *cli.Context) error {

	// Read cli arguments
	targetURL := c.String("targetURL")
	proxyAddrs := c.StringSlice("proxyAddr")
	proxyAddrFile := c.String("proxyAddrFile")

	// Load targets from file
	if len(proxyAddrFile) > 0 {
		log.Printf("read target file %s", proxyAddrFile)
		contents, err := ioutil.ReadFile(proxyAddrFile)
		if err != nil {
			return err
		}

		lines := strings.Split(strings.TrimSuffix(string(contents), "\n"), "\n")
		for _, line := range lines {
			proxyAddrs = append(proxyAddrs, line)
		}
	}

	// Start one thread proxy check
	for _, proxyAddr := range proxyAddrs {
		if err := check(targetURL, proxyAddr); err == nil {
			fmt.Printf("%s\n", proxyAddr)
		} else {
			log.Printf("invalid proxy %s: %s", proxyAddr, err)
		}
	}

	return nil
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
