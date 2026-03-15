# production-service

Production-oriented example that composes:

- API server on `:8080`
- metrics server on `:9090`
- background worker
- readiness drain mode during shutdown

Run:

```bash
go run ./examples/production-service
```

Try endpoints:

```bash
curl http://localhost:8080/
curl http://localhost:8080/healthz
curl http://localhost:8080/readyz
curl http://localhost:9090/metrics
```

Press `Ctrl+C` and quickly re-check readiness:

```bash
curl -i http://localhost:8080/readyz
```

Expected behavior during shutdown:

- API starts rejecting new traffic (`503`) while draining
- worker exits cooperatively
- process stops gracefully
