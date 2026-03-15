.PHONY: fmt test race vulncheck ci run-basic-example run-production-example

fmt:
	"$(shell go env GOROOT)/bin/gofmt" -w .

test:
	go test ./...

race:
	go test -race ./...

vulncheck:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

ci: test race vulncheck

run-basic-example:
	go run ./examples/basic-service

run-production-example:
	go run ./examples/production-service
