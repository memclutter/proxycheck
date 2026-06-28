package proxycheck

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveJudge(t *testing.T) {
	t.Run("azenv.php resolves", func(t *testing.T) {
		judge, err := ResolveJudge("azenv.php")
		assert.NoError(t, err)
		assert.IsType(t, AZEnvPhpJudge{}, judge)
	})

	t.Run("proxyjudge.us resolves", func(t *testing.T) {
		judge, err := ResolveJudge("proxyjudge.us")
		assert.NoError(t, err)
		assert.IsType(t, ProxyjudgeUsJudge{}, judge)
	})

	t.Run("unknown name errors", func(t *testing.T) {
		judge, err := ResolveJudge("foo")
		assert.Nil(t, judge)
		assert.Error(t, err)
		if err != nil {
			assert.Contains(t, err.Error(), "unknown judge: foo")
			assert.Contains(t, err.Error(), "azenv.php")
			assert.Contains(t, err.Error(), "proxyjudge.us")
		}
	})
}

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
