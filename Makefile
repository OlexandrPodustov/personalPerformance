lint: 
	which ${GOPATH}/bin/golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	${GOPATH}/bin/golangci-lint run

fmt:
	which ${GOPATH}/bin/gofumpt  || go install mvdan.cc/gofumpt@v0.4.0
	gofumpt -w .
	which ${GOPATH}/bin/gogroup  || go install github.com/vasi-stripe/gogroup/cmd/gogroup@v0.0.0-20200806161525-b5d7f67a97b5
	find . -type f -name '*.go' -exec gogroup -order std,other,prefix=github.com/OlexandrPodustov -rewrite {} \;
