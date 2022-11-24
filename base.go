package proxycheck

// Feed
// Data source interface. Encapsulates the logic of extracting the next proxy for verification.
type Feed interface {
	// Next
	// Returns a string in ip:port format of the next proxy to check
	Next() (string, error)
}
