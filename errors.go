package proxycheck

import "errors"

var (
	// FeedEnd this error happens when the feed runs out of proxies to check.
	// Named without the conventional Err* prefix to preserve the public API.
	//nolint:staticcheck // ST1012: renaming would break importers of the package
	FeedEnd = errors.New("feed end")
)
