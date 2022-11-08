package main

import (
	"log"
	"os"

	"github.com/memclutter/proxycheck"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "proxycheck"
	app.Description = "Proxy list checker"
	app.Version = "0.0.0"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "targetURL",
			Usage: "Target url for checking proxy",
			Value: "https://google.com",
		},
		&cli.StringSliceFlag{
			Name:  "proxyAddr",
			Usage: "Specify proxy address to check, format scheme://host:port",
		},
		&cli.StringFlag{
			Name:  "proxyAddrFile",
			Usage: "Specify proxy list file, contains on each line scheme://host:port",
		},
	}

	app.Action = proxycheck.Action

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("app run failed: %s", err)
	}
}
