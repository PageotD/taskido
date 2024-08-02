# Variables
GO_CMD = go
SRC_DIR = cmd/taskido
BIN_DIR = bin
INT_DIR = internal
INT_PKG = $(wildcard $(INT_DIR)/*)
BIN_NAME = taskido
PRF_DIR = profiling
BUILD_FLAGS = -ldflags="-s -w" -a -tags netgo -installsuffix netgo

# Targets
all: build

# Create necessary directories
mkdirs:
	mkdir -p $(BIN_DIR)

# Compile for specified OS and ARCH
build-for-os-arch:
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_CMD) build -C $(SRC_DIR) -o ../../$(BIN_DIR)/$(BIN_NAME)-$(GOOS)-$(GOARCH) $(BUILD_FLAGS)

# Compile for all platforms
build: mkdirs
	$(MAKE) build-for-os-arch GOOS=linux GOARCH=amd64
	$(MAKE) build-for-os-arch GOOS=windows GOARCH=amd64
	$(MAKE) build-for-os-arch GOOS=darwin GOARCH=amd64

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
	rm -f $(BIN_DIR)/*

.PHONY: all build clean test test-cmd test-internal
