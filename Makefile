fmt:
	@echo "==> Running gofmt..."
	gofmt -s -w .

build: fmt
	@echo "==> Building things..."
	go build -ldflags="-s -w" ./...

test:
	@echo "==> Running tests..."
	@go test -cover ./...

report:
	@echo "==> Generating report card..."
	@goreportcard-cli -v

bench: test
	@echo "==> Running benchmarks (may take a while)..."
	@go test -run=XXX -bench=. ./...

cover:
	@echo "==> Calculating coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out | grep -vE "^total" | sort -k3,3n
	@go tool cover -html=coverage.out

.PHONY: build fmt
