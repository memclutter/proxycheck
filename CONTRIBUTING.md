# Contributing

Thanks for your interest in improving `proxycheck`. This document describes how to
set up the project, the conventions the codebase follows, and how to get a change
merged.

## Prerequisites

- Go 1.18 or newer (the module targets `go 1.18`).
- [golangci-lint](https://golangci-lint.run/) for local linting (CI runs it too).
- No external services are required to build, but note that some checks make real
  network requests through proxy judges.

## Getting started

```bash
git clone https://github.com/memclutter/proxycheck.git
cd proxycheck
go mod download
go build ./cmd        # build the CLI
go test ./...         # run the suite
```

`proxycheck` is both a library (the repo-root package) and a binary
(`./cmd`). To try a change end to end, build the binary and feed it a proxy:

```bash
go build -o proxycheck ./cmd
echo "1.2.3.4:8080" | ./proxycheck
```

## Development workflow

- Branch off `main` with a short-lived topic branch; keep one logical change per
  pull request.
- Run the checks below before pushing; CI must be green before review.

```bash
go build ./...                              # compiles (library + cmd)
go test ./... -race -coverprofile=cover.out # runs the suite with the race detector
gofmt -l .                                  # must print nothing (formatting)
golangci-lint run                           # lints
```

Two GitHub Actions workflows run on every push and pull request:
`.github/workflows/go.yml` (build + `go test`, with coverage uploaded to Codecov)
and `.github/workflows/golangci-lint.yml` (linting). Keep both green.

## Code style

- Format with `gofmt` / `goimports`; do not hand-format.
- Wrap errors with context (`fmt.Errorf("...: %w", err)`) rather than returning
  bare errors where extra information helps.
- Prefer table-driven tests; the existing `*_test.go` files follow that pattern —
  add rows for the behaviour you change.
- The exported surface (`Check`, `CheckResult`, the `Feed` and `Judge`
  interfaces, their shipped implementations, the `Judges` registry, and
  `ProxyRequest`) is a stability contract. Treat changes to it deliberately.

## Adding a judge

Judges live in `judges.go`. To add one:

1. Implement the `Judge` interface (`TargetURL() string`, `Timeout()
   time.Duration`) on a new type.
2. Register it in the `Judges` map under a stable name.
3. Add a test, and mention the new judge in `README.md`.

## Adding a feed

Feeds live in `feed.go` and implement the `Feed` interface (`Next() (string,
error)`), returning the `FeedEnd` sentinel when exhausted. Add a test alongside
the existing `SliceFeed` / `FileFeed` cases.

## Commit messages

This project uses [Conventional Commits](https://www.conventionalcommits.org/):

```
feat(check): measure average speed across protocols
fix(feed): correct off-by-one in SliceFeed
docs(readme): document the output format
```

Use `feat` / `fix` / `docs` / `refactor` / `test` / `chore` as appropriate; mark
breaking changes with a `!` or a `BREAKING CHANGE:` footer.

## Pull requests

- Describe what changes and why; link any related issue.
- Keep the diff focused and the history readable.
- Update `README.md` and `CHANGELOG.md` (under `## [Unreleased]`) when your change
  affects behaviour users can see.

## Releases

Releases follow [Semantic Versioning](https://semver.org/). While the major
version is `0`, the public API may change in any release. Maintainers cut releases
by tagging `vMAJOR.MINOR.PATCH` and publishing notes built from the changelog.
