package proxycheck

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewSliceFeed(t *testing.T) {
	excepted := []string{"10.10.0.1:1002", "20.10.0.1:8888"}
	feed := NewSliceFeed(excepted)

	assert.Equal(t, 0, feed.idx, "must be start from first element")
	assert.Equal(t, excepted, feed.slice, "must be valid slice")
}

func TestSliceFeed_Next(t *testing.T) {
	slice := []string{"10.10.0.1:1002", "20.10.0.1:8888", "50.220.10.1:80", "30.20.10.1:999"}
	feed := NewSliceFeed(slice)

	for i, item := range slice {
		nextItem, nextErr := feed.Next()

		assert.NoError(t, nextErr, "must be without error")
		assert.Equal(t, item, nextItem, "must be valid next item from feed")
		assert.Equal(t, i+1, feed.idx, "must be correct next index")
	}

	// End test
	nextItem, nextErr := feed.Next()

	assert.EqualError(t, nextErr, FeedEnd.Error(), "must be return feed end error")
	assert.Equal(t, nextItem, "", "must be return empty on end")
}

func TestNewFileFeed(t *testing.T) {
	read := bytes.NewReader([]byte(`10.10.0.1:1002\n20.10.0.1:8888\n`))
	feed := NewFileFeed(read)

	assert.NotNil(t, feed.s, "must be set *io.Scanner")
}

func TestFileFeed_Next(t *testing.T) {
	slice := []string{"10.10.0.1:1002", "20.10.0.1:8888", "50.220.10.1:80", "30.20.10.1:999"}
	read := bytes.NewReader([]byte(strings.Join(slice, "\n")))
	feed := NewFileFeed(read)

	for _, item := range slice {
		nextItem, nextErr := feed.Next()

		assert.NoError(t, nextErr, "must be without errors")
		assert.Equal(t, item, nextItem, "must be valid next item from feed")
	}

	nextItem, nextErr := feed.Next()

	assert.EqualError(t, nextErr, FeedEnd.Error(), "must be return feed end error")
	assert.Equal(t, nextItem, "", "must be return empty on end")
}
