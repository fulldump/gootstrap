# Production Patterns

This guide shows practical patterns for running real services with `gootstrap`.

## 1) Compose API + worker + metrics

Use `RunAll` to coordinate all long-running components under one lifecycle.

```go
runner := gootstrap.RunAll(
    gootstrap.RunGracefulHttpServer(apiServer),
    gootstrap.RunHTTPServer(metricsServer),
    backgroundWorker,
)

gootstrap.RunUntilSignal(runner)
```

Why this matters:

- one process entrypoint
- one signal handling strategy
- one coordinated shutdown path

## 2) Readiness during shutdown

Expose `/readyz` and return `503` while draining new traffic.

```go
var draining atomic.Bool

mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
    if draining.Load() {
        w.WriteHeader(http.StatusServiceUnavailable)
        _, _ = w.Write([]byte("draining"))
        return
    }
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("ready"))
})
```

Set `draining = true` before invoking the API server stop function.

## 3) Liveness endpoint

Always expose an inexpensive liveness endpoint:

```go
mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("ok"))
})
```

## 4) Bounded, cooperative workers

Workers should support cooperative shutdown and avoid busy loops.

```go
stopCh := make(chan struct{})

start := func() error {
    t := time.NewTicker(2 * time.Second)
    defer t.Stop()
    for {
        select {
        case <-t.C:
            // do one unit of work
        case <-stopCh:
            return nil
        }
    }
}

stop := func() error {
    close(stopCh)
    return nil
}
```

## 5) Keep stop idempotent

Use `sync.Once` or equivalent for shutdown paths that may be called more than once.

## End-to-end example

See `examples/production-service/main.go`.

Run locally:

```bash
go run ./examples/production-service
```
