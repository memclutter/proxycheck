package proxycheck

import "time"

type Judge interface {
	TargetURL() string
	Timeout() time.Duration
}

type AZEnvPhpJudge struct {
}

func (j AZEnvPhpJudge) TargetURL() string      { return "http://www.wfuchs.de/azenv.php" }
func (j AZEnvPhpJudge) Timeout() time.Duration { return 500 * time.Millisecond }
