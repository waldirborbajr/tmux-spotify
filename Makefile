# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
OUTPUT_BIN=bin
BINARY_NAME=tmux-spotify-cli
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(OUTPUT_BIN)/$(BINARY_UNIX) -v -ldflags="-w -s" -trimpath

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(OUTPUT_BIN)/$(BINARY_NAME)
	rm -f $(OUTPUT_BIN)/$(BINARY_UNIX)

run:
	$(GOBUILD) -o $(OUTPUT_BIN)/$(BINARY_NAME) -v ./...
	./$(OUTPUT_BIN)/$(BINARY_NAME)

deps:
	$(GOGET) github.com/joho/godotenv
	$(GOGET) github.com/zmb3/spotify
	$(GOGET) golang.org/x/oauth2

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(OUTPUT_BIN)/$(BINARY_UNIX) -v -ldflags="-w -s" -trimpath

# Build for development with race detection
dev:
	$(GOBUILD) -race -o $(OUTPUT_BIN)/$(BINARY_NAME) -v ./...

.PHONY: all build test clean run deps build-linux dev
