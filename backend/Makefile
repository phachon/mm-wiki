GO = go
APP_NAME = mm-wiki
ENV ?= dev
CONF_PATH = ./conf/$(ENV)/app.yaml

.PHONY: all build run dev test clean lint linux windows darwin help

all: build

build:
	$(GO) build -o $(APP_NAME) .

run: build
	./$(APP_NAME) -conf $(CONF_PATH)

dev:
	$(GO) run . -conf $(CONF_PATH)

test:
	$(GO) test $(M) -v -gcflags=all=-l -coverpkg=./... -coverprofile=test.out ./...

lint:
	@if command -v golangci-lint > /dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed, running go vet instead"; \
		$(GO) vet ./...; \
	fi

clean:
	rm -f $(APP_NAME) test.out
	rm -rf release release-*

linux:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(APP_NAME) .

windows:
	GOOS=windows GOARCH=amd64 $(GO) build -o $(APP_NAME).exe .

darwin:
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(APP_NAME) .

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build     Build the backend binary"
	@echo "  run       Build and run the backend server (ENV=dev|test|prod)"
	@echo "  dev       Run with go run (hot reload friendly, ENV=dev|test|prod)"
	@echo "  test      Run all tests"
	@echo "  lint      Run linter (golangci-lint or go vet)"
	@echo "  clean     Remove build artifacts"
	@echo "  linux     Cross-compile for Linux amd64"
	@echo "  windows   Cross-compile for Windows amd64"
	@echo "  darwin    Cross-compile for macOS amd64"
	@echo "  help      Show this help message"