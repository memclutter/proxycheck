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
		Version:     "0.0.1",
		Flags: []cli.Flag{
			// @TODO replace with judge name
			//&cli.StringFlag{
			//	Name:  "targetURL",
			//	Usage: "Target url for checking proxy",
			//	Value: "https://google.com",
			//},
			&cli.IntFlag{Name: "threads", Usage: "Count of threads", Value: 10},
		},
		Action: proxycheck.Action,
	}).Run(os.Args); err != nil {
		log.Fatalf("app run failed: %s", err)
	}
}
