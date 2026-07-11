GOLANGCI_LINT_VERSION := v2.12.2

all: test fmt

lint:
	which ${GOPATH}/bin/golangci-lint || go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}
	${GOPATH}/bin/golangci-lint run

fmt:
	which ${GOPATH}/bin/golangci-lint || go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}
	${GOPATH}/bin/golangci-lint fmt

test:
	go test ./...

vulncheck:
	govulncheck ./...
