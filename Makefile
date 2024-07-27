# Makefile for the API server
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY_NAME=nerdstore
BINARY_UNIX=$(BINARY_NAME)_unix
# Build parameters
BUILD_DIR=./bin
CMD_DIR=./cmd/api
# Colors for terminal output
GREEN=\033[0;32m
NC=\033[0m
.PHONY: all build clean test run deps

all: test build

build:
	@echo "$(GREEN)Building...$(NC)"
	$(GOBUILD) -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

clean:
	@echo "$(GREEN)Cleaning...$(NC)"
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

test:
	@echo "$(GREEN)Testing...$(NC)"
	$(GOTEST) -v ./...

run:
	@echo "$(GREEN)Running API server...$(NC)"
	$(GOCMD) run $(CMD_DIR)

deps:
	@echo "$(GREEN)Downloading dependencies...$(NC)"
	$(GOMOD) download

seed:
	@echo "$(GREEN)Seeding database...$(NC)"
	$(GOCMD) run ./scripts/seed.go


# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_UNIX) $(CMD_DIR)

# Help target
help:
	@echo "$(GREEN)Available targets:$(NC)"
	@echo "  all         - Run tests and build"
	@echo "  build       - Build the application"
	@echo "  clean       - Clean build artifacts"
	@echo "  test        - Run tests"
	@echo "  run         - Run the API server"
	@echo "  deps        - Download dependencies"
	@echo "  docker-up   - Start Docker containers"
	@echo "  docker-down - Stop Docker containers"
	@echo "  build-linux - Build for Linux (cross-compilation)"