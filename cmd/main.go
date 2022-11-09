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
			&cli.StringSliceFlag{
				Name:  "proxyAddr",
				Usage: "Specify proxy address to check, format scheme://host:port",
			},
			&cli.StringFlag{
				Name:  "proxyAddrFile",
				Usage: "Specify proxy list file, contains on each line scheme://host:port",
			},
		},
		Action: proxycheck.Action,
	}).Run(os.Args); err != nil {
		log.Fatalf("app run failed: %s", err)
	}
}
