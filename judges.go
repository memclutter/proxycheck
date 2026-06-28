package proxycheck

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

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

// ResolveJudge returns the Judge registered under name. An unknown name yields
// an error that lists the available judge names.
func ResolveJudge(name string) (Judge, error) {
	judge, ok := Judges[name]
	if !ok {
		return nil, fmt.Errorf("unknown judge: %s (available: %s)", name, judgeNames())
	}
	return judge, nil
}

// judgeNames returns the registered judge names, sorted and comma-joined, so
// error messages are deterministic.
func judgeNames() string {
	names := make([]string, 0, len(Judges))
	for name := range Judges {
		names = append(names, name)
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}
