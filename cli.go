package proxycheck

import (
	"fmt"
	cli "github.com/urfave/cli/v2"
	"os"
	"strings"
	"sync"
)

func Action(c *cli.Context) error {

	// If pass args, use it or read stdin as input file with proxies
	var feed Feed

	if c.Args().Len() > 0 {
		feed = NewSliceFeed(c.Args().Slice())
	} else {
		feed = NewFileFeed(os.Stdin)
	}

	// Worker pool
	poolSize := c.Int("threads")
	proxyAddrs := make(chan string, poolSize)
	var wg sync.WaitGroup
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go func(proxyAddrs chan string) {
			defer wg.Done()
			for proxyAddr := range proxyAddrs {
				if res := Check(proxyAddr, &AZEnvPhpJudge{}); res.Online {
					fmt.Printf("%s\t%s\t%s\n", proxyAddr, strings.Join(res.Protocols, ","), res.Speed.String())
				} else {
					fmt.Fprintf(os.Stderr, "invalid proxy %s: %v\n", proxyAddr, res.Err)
				}
			}
		}(proxyAddrs)
	}

	// Start one thread proxy check
	for {
		proxyAddr, err := feed.Next()
		if err == FeedEnd {
			break
		}

		proxyAddrs <- proxyAddr
	}

	wg.Wait()
	return nil
}
