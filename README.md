# proxycheck

> Proxy list checker for Go — feed it a list of `ip:port` proxies and it reports
> which are alive, which protocols (`http`, `https`, `socks4`, `socks5`) each
> supports, and how fast each responds.

[![Release](https://img.shields.io/github/v/release/memclutter/proxycheck?sort=semver)](https://github.com/memclutter/proxycheck/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/memclutter/proxycheck.svg)](https://pkg.go.dev/github.com/memclutter/proxycheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/memclutter/proxycheck)](https://goreportcard.com/report/github.com/memclutter/proxycheck)
[![Go version](https://img.shields.io/github/go-mod/go-version/memclutter/proxycheck)](go.mod)
[![CI](https://github.com/memclutter/proxycheck/actions/workflows/go.yml/badge.svg)](https://github.com/memclutter/proxycheck/actions/workflows/go.yml)
[![golangci-lint](https://github.com/memclutter/proxycheck/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/memclutter/proxycheck/actions/workflows/golangci-lint.yml)
[![codecov](https://codecov.io/gh/memclutter/proxycheck/branch/main/graph/badge.svg?token=PFJO0VOY09)](https://codecov.io/gh/memclutter/proxycheck)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

`proxycheck` validates raw proxy lists. Each proxy is tested by routing a request
to an external **proxy judge** through it: if the judge answers `HTTP 200`, the
proxy is considered to work for that protocol. The tool tries all four protocols
per address and reports the ones that succeed, together with a response-speed
figure.

It ships two ways:

- a **CLI binary** (`proxycheck`) that reads proxies from arguments or stdin and
  prints the working ones, and
- an **importable Go package** (`github.com/memclutter/proxycheck`) exposing the
  `Check` function and the `Feed` / `Judge` interfaces for use in other programs.

## Contents

- [Install](#install)
- [CLI usage](#cli-usage)
- [Output format](#output-format)
- [Library usage](#library-usage)
- [How it works](#how-it-works)
- [Contributing](#contributing)
- [License](#license)

## Install

Build the binary from source (Go 1.18+):

```shell
go install github.com/memclutter/proxycheck/cmd@latest
```

or clone and build locally:

```shell
git clone https://github.com/memclutter/proxycheck.git
cd proxycheck
go build -o proxycheck ./cmd
```

## CLI usage

Pass proxies as arguments:

```shell
proxycheck 108.20.30.1:8080 89.33.123.100:3128
```

or pipe a list (one `ip:port` per line, blank lines ignored) on stdin:

```shell
cat proxies.txt | proxycheck --threads 50
```

| Flag        | Default        | Description                                  |
|-------------|----------------|----------------------------------------------|
| `--threads` | `10`           | Number of proxies checked concurrently.      |
| `--judge`   | `proxyjudge.us`| Proxy judge name (see [How it works](#how-it-works)). |

An address may be a bare `ip`, in which case port `80` is assumed.

## Output format

Working proxies are written to **stdout**, one per line, tab-separated:

```
<addr>\t<protocols>\t<speed>
```

- `<protocols>` — a comma-separated subset of `http,https,socks4,socks5`.
- `<speed>` — a Go duration string, e.g. `412ms`.

Proxies that fail every protocol are reported on **stderr** and are absent from
stdout, so you can pipe the clean list onward while still seeing failures:

```shell
cat proxies.txt | proxycheck > working.tsv 2> failures.log
```

## Library usage

```go
package main

import (
	"fmt"

	"github.com/memclutter/proxycheck"
)

func main() {
	res := proxycheck.Check("108.20.30.1:8080", proxycheck.AZEnvPhpJudge{})
	if res.Online {
		fmt.Printf("online via %v in %s\n", res.Protocols, res.Speed)
	} else {
		fmt.Printf("offline: %v\n", res.Err)
	}
}
```

The package also exports the `Feed` interface (`SliceFeed`, `FileFeed`) for
streaming proxy sources, the `Judge` interface with the shipped judges and the
`Judges` registry, and `ProxyRequest` for a single timed request through a proxy.

## How it works

A **judge** is an external endpoint that echoes request details; reaching it
through a proxy proves the proxy forwards traffic. Two judges ship:

| Name            | Target URL                          | Timeout |
|-----------------|-------------------------------------|---------|
| `azenv.php`     | `http://www.wfuchs.de/azenv.php`    | 3s      |
| `proxyjudge.us` | `http://proxyjudge.us/`             | 1s      |

For each proxy, `Check` tries `http`, `https`, `socks4`, and `socks5` in turn
(HTTP/HTTPS via Go's proxy transport, SOCKS via `h12.io/socks`). A protocol
counts as working only when the request completes with an `HTTP 200`. The proxy
is "online" if at least one protocol succeeds.

## Contributing

Contributions are welcome — see [CONTRIBUTING.md](CONTRIBUTING.md) for setup,
coding conventions, and the commit/PR process. Changes are recorded in
[CHANGELOG.md](CHANGELOG.md).

## License

Released under the [MIT License](LICENSE).
