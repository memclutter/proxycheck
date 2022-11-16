package proxycheck

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func Action(c *cli.Context) error {

	// If pass args, use it or read stdin as input file with proxies
	var feed Feed

	if c.Args().Len() > 0 {
		feed = NewSliceFeed(c.Args().Slice())
	} else {
		feed = NewFileFeed(os.Stdin)
	}

	// Start one thread proxy check
	for {
		proxyAddr, err := feed.Next()
		if err == FeedEnd {
			break
		}

		if res := Check(proxyAddr, &AZEnvPhpJudge{}); res.Online {
			fmt.Printf("%s\t%s\t%s\n", proxyAddr, strings.Join(res.Protocols, ","), res.Speed.String())
		} else {
			fmt.Fprintf(os.Stderr, "invalid proxy %s: %v\n", proxyAddr, res.Err)
		}
	}

	return nil
}
