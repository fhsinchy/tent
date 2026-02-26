VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo development)
MODULE  := github.com/fhsinchy/tent
BINARY  := tent
BINDIR  := bin

LDFLAGS := -ldflags="-X '$(MODULE)/cmd.version=$(VERSION)'"

.PHONY: build install clean vet lint fmt check test

## Build

build:
	go build $(LDFLAGS) -o $(BINDIR)/$(BINARY)

install:
	go install $(LDFLAGS)

clean:
	rm -rf $(BINDIR)

## Quality

vet:
	go vet ./...

fmt:
	gofmt -s -w .

lint: vet
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed, skipping"; \
	fi

test:
	go test ./...

check: fmt vet lint test

## Helpers

tidy:
	go mod tidy

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Build:"
	@echo "  build     Build binary to $(BINDIR)/$(BINARY)"
	@echo "  install   Install binary via go install"
	@echo "  clean     Remove build artifacts"
	@echo ""
	@echo "Quality:"
	@echo "  vet       Run go vet"
	@echo "  fmt       Format code with gofmt"
	@echo "  lint      Run vet + golangci-lint (if installed)"
	@echo "  test      Run tests"
	@echo "  check     Run fmt + vet + lint + test"
	@echo ""
	@echo "Helpers:"
	@echo "  tidy      Run go mod tidy"
	@echo "  help      Show this help"
