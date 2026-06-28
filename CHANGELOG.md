# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
While the major version is `0`, the public API may change in any release.

## [Unreleased]

## [0.0.6] - 2026-06-28

### Added

- `CONTRIBUTING.md` describing setup, the test/lint workflow, code style, and the
  commit/PR/release process.
- This `CHANGELOG.md`, with reconstructed notes for the previously empty
  `v0.0.1`–`v0.0.5` releases.
- `ResolveJudge` helper that looks a judge up in the `Judges` registry, and a
  `NewApp` constructor shared by the binary and the tests.

### Changed

- Relicensed the project from the Apache License 2.0 to the MIT License, matching
  the other `memclutter` projects.
- Reworked `README.md`: a description, status badges, install/usage sections, the
  tab-separated output format, a library-usage example, and a "how it works"
  explanation of judges and the protocol probing.

### Fixed

- `--judge` now selects the proxy judge. It was parsed but ignored: every check
  ran against the AZEnv judge regardless of the flag, so the documented
  `proxyjudge.us` default never applied. An unknown judge name is now rejected
  with a non-zero exit, listing the valid names, and no proxies are checked.
- The CLI no longer hangs after processing all proxies. `Action` never closed its
  worker channel, so the worker goroutines never returned and the program blocked
  forever; the channel is now closed once the feed is exhausted.
- Rewrote the GitHub Actions workflows: pinned current Go versions, fixed the
  empty `go-version` in the lint job, and dropped the dead OS matrix.

## [0.0.5] - 2022-11-25

### Added

- Second proxy judge `ProxyjudgeUsJudge` (`http://proxyjudge.us/`) alongside
  `AZEnvPhpJudge`, plus a `Judges` registry map keyed by judge name.
- `golangci-lint` and its GitHub Actions workflow.
- More tests around the judges and check engine.

## [0.0.4] - 2022-11-25

### Added

- Tests for the `Feed` implementations.

### Fixed

- Off-by-one bug in `SliceFeed.Next` that mis-tracked the current index.

## [0.0.3] - 2022-11-24

### Changed

- Introduced `base.go` holding the core `Feed` and `Judge` interfaces (moved out
  of their previous files) and reformatted the package godoc.

## [0.0.2] - 2022-11-17

A full rewrite of the checker into a concurrent, judge-based engine.

### Added

- `Check(addr, judge)` engine that probes each `ip:port` against all four
  protocols (`http`, `https`, `socks4`, `socks5`) and returns a `CheckResult`
  with the online flag, supported protocols, per-protocol errors, and speed.
- `Judge` abstraction (target URL + recommended timeout) with the `AZEnvPhpJudge`
  implementation.
- Reading proxies from a stdin stream in addition to command-line arguments.
- Concurrent checking via a worker pool sized by the `--threads` flag
  (default 10).
- GitHub Actions CI running the test suite with coverage uploaded to Codecov.

### Changed

- Upgraded the CLI from `urfave/cli` v1 to v2 and moved the entry point to
  `cmd/main.go`.
- Proxy addresses are now plain `ip:port` (the tool tries every protocol) instead
  of requiring a `scheme://` prefix.
- Output is now tab-separated `addr<TAB>protocols<TAB>speed` for working proxies.

## [0.0.1] - 2019-03-21

First release of `proxycheck` — a simple proxy list checker.

### Added

- A `urfave/cli` v1 command that checks proxies supplied via repeatable
  `--proxyAddr` flags or a `--proxyAddrFile` list file.
- Per-scheme checking for `http`, `https`, `socks4`, and `socks5` proxies
  (`h12.io/socks` for SOCKS), validating each against a `--targetURL`.
- Working proxy addresses printed to stdout; failures logged.

[Unreleased]: https://github.com/memclutter/proxycheck/compare/v0.0.5...HEAD
[0.0.5]: https://github.com/memclutter/proxycheck/compare/v0.0.4...v0.0.5
[0.0.4]: https://github.com/memclutter/proxycheck/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/memclutter/proxycheck/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/memclutter/proxycheck/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/memclutter/proxycheck/releases/tag/v0.0.1
