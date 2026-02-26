VERSION    := $(shell git describe --tags --always --dirty 2>/dev/null || echo development)
MODULE     := github.com/fhsinchy/tent
BINARY     := tent
BINDIR     := bin
BUILDTAGS  := containers_image_openpgp

GOFLAGS := CGO_ENABLED=0
LDFLAGS := -ldflags="-s -w -X '$(MODULE)/cmd.version=$(VERSION)'"
TAGS    := -tags $(BUILDTAGS)

.PHONY: build install clean vet lint fmt check test

## Build

build:
	$(GOFLAGS) go build $(TAGS) $(LDFLAGS) -o $(BINDIR)/$(BINARY)

install:
	$(GOFLAGS) go install $(TAGS) $(LDFLAGS)

clean:
	rm -rf $(BINDIR)

## Quality

vet:
	$(GOFLAGS) go vet $(TAGS) ./...

fmt:
	gofmt -s -w .

lint: vet
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed, skipping"; \
	fi

test:
	$(GOFLAGS) go test $(TAGS) ./...

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
