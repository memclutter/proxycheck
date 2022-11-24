package proxycheck

import (
	"bufio"
	"os"
)

// SliceFeed
// Proxy feed from string slice like this []string{"108.20.30.1:500", "89.33.123.100:40", "50.73.100.1:55"}
type SliceFeed struct {
	idx   int
	slice []string
}

func NewSliceFeed(slice []string) *SliceFeed { return &SliceFeed{idx: 0, slice: slice} }

func (f *SliceFeed) Next() (string, error) {
	if f.idx < len(f.slice)-1 {
		f.idx += 1
		return f.slice[f.idx-1], nil
	} else {
		return "", FeedEnd
	}
}

// FileFeed
// Proxy feed from file stream, like stdin, os file.
type FileFeed struct {
	s *bufio.Scanner
}

func NewFileFeed(file *os.File) *FileFeed {
	feed := new(FileFeed)
	feed.s = bufio.NewScanner(file)
	return feed
}

func (f FileFeed) Next() (string, error) {
	if f.s.Scan() {
		return f.s.Text(), nil
	} else if err := f.s.Err(); err != nil {
		return "", err
	} else {
		return "", FeedEnd
	}
}
