package main

import (
	"log"
	"os"

	"github.com/memclutter/proxycheck"
	"github.com/urfave/cli/v2"
)

func main() {
	if err := (&cli.App{
		Name:        "proxycheck",
		Description: "Proxy checker tool",
		Version:     "0.0.4",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "judge", Usage: "Set judge", Value: "proxyjudge.us"},
			&cli.IntFlag{Name: "threads", Usage: "Count of threads", Value: 10},
		},
		Action: proxycheck.Action,
	}).Run(os.Args); err != nil {
		log.Fatalf("app run failed: %s", err)
	}
}
