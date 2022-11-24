package proxycheck

import "time"

type AZEnvPhpJudge struct{}

func (j AZEnvPhpJudge) TargetURL() string      { return "http://www.wfuchs.de/azenv.php" }
func (j AZEnvPhpJudge) Timeout() time.Duration { return 3 * time.Second }

type ProxyjudgeUsJudge struct{}

func (j ProxyjudgeUsJudge) TargetURL() string      { return "http://proxyjudge.us/" }
func (j ProxyjudgeUsJudge) Timeout() time.Duration { return 1 * time.Second }

var Judges = map[string]Judge{
	"azenv.php":     AZEnvPhpJudge{},
	"proxyjudge.us": ProxyjudgeUsJudge{},
}
