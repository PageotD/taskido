# Variables
GO_CMD = go
SRC_DIR = cmd/taskido
BIN_DIR = bin
INT_DIR = internal
INT_PKG = $(wildcard $(INT_DIR)/*)
BIN_NAME = taskido
PRF_DIR = profiling

# Targets
all: build

# Compile for Linux
build-linux:
	GOOS=linux GOARCH=amd64 $(GO_CMD) build -C $(SRC_DIR) -o ../../$(BIN_DIR)/$(BIN_NAME)-linux -ldflags="-s -w" -a -tags netgo -installsuffix netgo 

# Compile for Windows
build-windows:
	GOOS=windows GOARCH=amd64 $(GO_CMD) build -C $(SRC_DIR) -o ../../$(BIN_DIR)/$(BIN_NAME).exe -ldflags="-s -w" -a -tags netgo -installsuffix netgo

# Compile for macOS
build-macos:
	GOOS=darwin GOARCH=amd64 $(GO_CMD) build -C $(SRC_DIR) -o ../../$(BIN_DIR)/$(BIN_NAME)-macos -ldflags="-s -w" -a -tags netgo -installsuffix netgo

# Build for all platforms
build: build-linux build-windows build-macos

# Test the cmd/taskido package
test-cmd:
	$(GO_CMD) -C $(SRC_DIR) test -v -coverprofile=../../$(PRF_DIR)/coverage-cmd.out -cpuprofile=../../$(PRF_DIR)/cpu-cmd.out -memprofile=../../$(PRF_DIR)/mem-cmd.out 

# Test the internal package
test-internal:
	for pkg in $(INT_PKG); do\
		$(GO_CMD) -C $${pkg} test -v -coverprofile=../../$(PRF_DIR)/coverage-internal.out -cpuprofile=../../$(PRF_DIR)/cpu-internal.out -memprofile=../../$(PRF_DIR)/mem-internal.out;\
	done

# Run all tests
test: test-cmd test-internal

# Clean build artifacts
clean:
	rm -f ../$(BIN_DIR)/$(BIN_NAME)*

.PHONY: all build clean test test-cmd test-internal
