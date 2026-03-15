.PHONY: test race vulncheck ci

test:
	go test ./...

race:
	go test -race ./...

vulncheck:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

ci: test race vulncheck
