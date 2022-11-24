package proxycheck

import "time"

// Feed
// Data source interface. Encapsulates the logic of extracting the next proxy for verification.
type Feed interface {
	// Next
	// Returns a string in ip:port format of the next proxy to check
	Next() (string, error)
}

// Judge
// The proxy judge encapsulates the logic of interaction with the proxy judge
type Judge interface {
	// TargetURL Returns the url to which you need to contact through the proxy
	TargetURL() string
	// Timeout Returns the recommended timeout for this referee
	Timeout() time.Duration
}
