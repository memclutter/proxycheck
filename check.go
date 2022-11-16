package proxycheck

import (
	"fmt"
	"net/url"
	"time"
)

type CheckResult struct {
	Protocols []string
	Online    bool
	Err       map[string]error
	Speed     time.Duration
}

func Check(proxyAddr string, judge Judge) (result CheckResult) {
	result.Err = make(map[string]error)

	ip, port, err := proxyAddrToIpPort(proxyAddr)
	if err != nil {
		result.Err[""] = fmt.Errorf("parse addr error: %v", err)
		return
	}

	result.Err = make(map[string]error)
	for _, protocol := range []string{"http", "https", "socks4", "socks5"} {
		proxyURL, err := url.Parse(fmt.Sprintf("%s://%s:%d", protocol, ip, port))
		if err != nil {
			result.Err[protocol] = fmt.Errorf("could not parse proxy address as proxyURL: %v", err)
			return
		}

		cpResult := checkProtocol(proxyURL, judge)
		if cpResult.Online {
			result.Protocols = append(result.Protocols, protocol)
			result.Online = true
			// @TODO average speed
			if result.Speed < cpResult.Speed {
				result.Speed = cpResult.Speed
			}
		} else {
			result.Err[protocol] = cpResult.Err
		}
	}
	return
}

type checkProtocolResult struct {
	Online bool
	Err    error
	Speed  time.Duration
}

func checkProtocol(proxyURL *url.URL, judge Judge) (result checkProtocolResult) {
	speedStart := time.Now().UTC()
	if _, err := ProxyRequest(judge.TargetURL(), proxyURL, judge.Timeout()); err == nil {
		result.Online = true
		speed := time.Now().UTC().Sub(speedStart)
		if result.Speed < speed {
			result.Speed = speed
		}
	} else {
		result.Err = err
	}
	return
}
