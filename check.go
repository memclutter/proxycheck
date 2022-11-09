package proxycheck

import (
	"net"
	"time"
)

type CheckResult struct {
	Protocols []string
	Online    bool
	Err       error
	Speed     time.Duration
}

func Check(ip net.IP, port int64, judge *Judge) (result CheckResult) {
	return
}
