package proxycheck

import "errors"

var (
	// FeedEnd this error happens when the feed runs out of proxies to check
	FeedEnd = errors.New("feed end")
)
