# basic-service

Simple executable example that shows how to compose multiple runners.

It starts:

- HTTP server on `:8080`
- background worker that logs every 2 seconds

Run:

```bash
go run ./examples/basic-service
```

Test endpoint:

```bash
curl http://localhost:8080
```

Stop with `Ctrl+C` to trigger graceful shutdown.
