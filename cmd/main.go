package main

import (
	"log"
	"os"

	"github.com/memclutter/proxycheck"
)

func main() {
	if err := proxycheck.NewApp().Run(os.Args); err != nil {
		log.Fatalf("app run failed: %s", err)
	}
}
