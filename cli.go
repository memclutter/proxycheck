package proxycheck

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"strings"
)

func Action(c *cli.Context) error {

	// Read cli arguments
	//targetURL := c.String("targetURL")
	proxyAddrs := c.StringSlice("proxyAddr")
	proxyAddrFile := c.String("proxyAddrFile")

	// Load targets from file
	if len(proxyAddrFile) > 0 {
		log.Printf("read target file %s", proxyAddrFile)
		contents, err := ioutil.ReadFile(proxyAddrFile)
		if err != nil {
			return err
		}

		lines := strings.Split(strings.TrimSuffix(string(contents), "\n"), "\n")
		for _, line := range lines {
			proxyAddrs = append(proxyAddrs, line)
		}
	}

	// Start one thread proxy check
	for _, proxyAddr := range proxyAddrs {
		if res := Check(proxyAddr, &AZEnvPhpJudge{}); res.Online {
			fmt.Printf("%s\t%s\t%s\n", proxyAddr, strings.Join(res.Protocols, ","), res.Speed.String())
		} else {
			fmt.Printf("invalid proxy %s: %v\n", proxyAddr, res.Err)
		}
	}

	return nil
}
