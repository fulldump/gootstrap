# gootstrap

[![CI](https://github.com/fulldump/gootstrap/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/fulldump/gootstrap/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/fulldump/gootstrap.svg)](https://pkg.go.dev/github.com/fulldump/gootstrap)
[![Go Report Card](https://goreportcard.com/badge/github.com/fulldump/gootstrap)](https://goreportcard.com/report/github.com/fulldump/gootstrap)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Lightweight lifecycle helpers to bootstrap long-running Go services.

`gootstrap` uses a tiny abstraction (`Runner`) and a small set of composable
helpers to manage process startup, coordinated shutdown, and OS signal handling.

## Install

```bash
go get github.com/fulldump/gootstrap
```

## Quick Start

```go
package main

import (
	"net/http"

	"github.com/fulldump/gootstrap"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("ok"))
		}),
	}

	gootstrap.Run(gootstrap.RunHTTPServer(server))
}
```

## Core Concepts

`Runner` defines lifecycle in two functions:

- `start() error` starts a service and blocks while it runs
- `stop() error` gracefully stops it

Helpers:

- `Run(...)` starts services and stops on `SIGINT`/`SIGTERM`
- `RunAll(...)` composes multiple runners
- `RunUntilSignal(...)` allows custom signals
- `RunHTTPServer(...)` wraps a standard `net/http` server
- `RunGracefulHttpServer(...)` adds a graceful-drain window

## Examples

Basic service (`examples/basic-service/main.go`) runs:

- an HTTP server
- a background worker
- signal-based graceful shutdown

Production-oriented service (`examples/production-service/main.go`) runs:

- API server + readiness/liveness endpoints
- separate metrics server
- background worker with cooperative shutdown
- drain mode before graceful stop

Run examples locally:

```bash
go run ./examples/basic-service
go run ./examples/production-service
```

## Production Guide

See `PRODUCTION_PATTERNS.md` for practical patterns used in real deployments.

## Stability and Versioning

The project follows semantic versioning. Release notes are tracked in
`CHANGELOG.md`.

Release process is documented in `RELEASING.md`.

## Contributing

- `CONTRIBUTING.md`
- `CODE_OF_CONDUCT.md`
- `SECURITY.md`
- `SUPPORT.md`

## License

MIT, see `LICENSE`.
