package proxycheck

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestJudge_TargetURL(t *testing.T) {
	testingTable := []struct {
		title string
		judge Judge
	}{
		{
			"correct azenvphp judge",
			AZEnvPhpJudge{},
		},
		{
			"correct proxyjudge.us",
			ProxyjudgeUsJudge{},
		},
	}

	for _, table := range testingTable {
		t.Run(table.title, func(t *testing.T) {
			targetURL := table.judge.TargetURL()
			assert.NotEmpty(t, targetURL, "must be not empty target url")
			_, err := url.Parse(targetURL)
			assert.NoError(t, err, "must be valid url address")
		})
	}
}

func TestJudge_Timeout(t *testing.T) {
	testingTable := []struct {
		title string
		judge Judge
	}{
		{
			"correct azenvphp judge",
			AZEnvPhpJudge{},
		},
		{
			"correct proxyjudge.us",
			ProxyjudgeUsJudge{},
		},
	}

	for _, table := range testingTable {
		t.Run(table.title, func(t *testing.T) {
			timeout := table.judge.Timeout()
			assert.NotEmpty(t, timeout, "must be not empty timeout")
		})
	}
}
