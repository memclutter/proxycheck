package proxycheck

import (
	"fmt"
	cli "github.com/urfave/cli/v2"
	"os"
	"strings"
	"sync"
)

// NewApp builds the proxycheck CLI application. It is shared by the cmd binary
// and the tests so both exercise the same flag wiring.
func NewApp() *cli.App {
	return &cli.App{
		Name:        "proxycheck",
		Description: "Proxy checker tool",
		Version:     "0.0.5",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "judge", Usage: "Set judge", Value: "proxyjudge.us"},
			&cli.IntFlag{Name: "threads", Usage: "Count of threads", Value: 10},
		},
		Action: Action,
	}
}

func Action(c *cli.Context) error {

	// Resolve the judge from the --judge flag before doing any work; an unknown
	// name aborts here, so no proxy is checked.
	judge, err := ResolveJudge(c.String("judge"))
	if err != nil {
		return err
	}

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
				if res := Check(proxyAddr, judge); res.Online {
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

	// Closing the channel lets the workers' range loops finish so wg.Wait
	// returns once the feed is exhausted.
	close(proxyAddrs)

	wg.Wait()
	return nil
}
