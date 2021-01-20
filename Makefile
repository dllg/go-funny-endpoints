.TARGET=build
.PHONY: clean format lint pre-build build build-linux build-osx test it docker-builder docker-build

BUILD_COMMIT=$(shell [ -z "$(SHA_DOCKERFILE)" ] && git log --pretty=format:'%h' -n 1 || echo "$(SHA_DOCKERFILE)")
BUILD_DATE=$(shell TZ=Utc date +%Y-%m-%d)
BUILD_TIME=$(shell TZ=Utc date +%H:%M:%S)
BUILD_DIR=build
BUILD_RELEASE=$(shell [ -z "$(TAG_DOCKERFILE)" ] && git describe --tags --exact-match 2> /dev/null || echo "$(TAG_DOCKERFILE)")
GIT_HASH=$(shell git rev-parse HEAD)

GO_LDFLAGS= -ldflags="-X $(PROJECT)/version.Commit=$(BUILD_COMMIT) -X $(PROJECT)/version.Date=$(BUILD_DATE) -X $(PROJECT)/version.Release=$(BUILD_RELEASE) -X $(PROJECT)/version.Time=$(BUILD_TIME)"
GO_BUILD_COMMAND=go build $(GO_LDFLAGS)
PWD=$(shell pwd)
GO_NAME=$(shell basename "$(PWD)")
PROJECT=github.com/dllg/$(GO_NAME)

clean:
	@echo "Cleaning..."
	@rm -Rf $(BUILD_DIR)

deps:
	go get -u golang.org/x/lint/golint

format:
	@echo "Automatically formating all your Go source code..."
	go fmt ./...


genmocks:
	@echo "Generating mocks"
	$(shell mockgen -source=httpclient/httpclient.go -destination=httpclient/mockhttpclient.go -package=httpclient)

# Lint the go code. Note: golint doesn't support vendor folder exclusion so we use find to filter it out
lint:
	@echo "Using vet to check for common mistakes..."
	@go vet ./...
	@echo "Checking style with golint..."
	@find . -type d -not -path "./vendor*" -exec golint {} \;

test:
	@echo "Running all unit tests..."
	go test ./...

pre-build: genmocks lint test
	@mkdir -p $(BUILD_DIR)

# Build using your computer's architecture
build: pre-build
	@echo "Building..."
	$(GO_BUILD_COMMAND) -o $(BUILD_DIR)/$(GO_NAME) -v

# Cross compile for linux
build-linux: pre-build
	@echo "Building Linux binary..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD_COMMAND) -o $(BUILD_DIR)/$(GO_NAME)-linux

docker-build:
	docker build --force-rm -t $(GO_NAME) .
