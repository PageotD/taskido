# Variables
GO_CMD = go
SRC_DIR = sources
BIN_DIR = bin
BIN_NAME = taskido
PRF_DIR = profiling

# Targets
all: build

# Compile for Linux
build-linux:
	GOOS=linux GOARCH=amd64 $(GO_CMD) -C $(SRC_DIR) build -o ../$(BIN_DIR)/$(BIN_NAME)-linux -ldflags="-s -w" -a -tags netgo -installsuffix netgo 

# Compile for Windows
build-windows:
	GOOS=windows GOARCH=amd64 $(GO_CMD) -C $(SRC_DIR) build -o ../$(BIN_DIR)/$(BIN_NAME).exe -ldflags="-s -w" -a -tags netgo -installsuffix netgo

# Compile for macOS
build-macos:
	GOOS=darwin GOARCH=amd64 $(GO_CMD) -C $(SRC_DIR) build -o ../$(BIN_DIR)/$(BIN_NAME)-macos -ldflags="-s -w" -a -tags netgo -installsuffix netgo

# Build for both platforms
build: build-linux build-windows build-macos

test:
	$(GO_CMD) -C $(SRC_DIR) test -v -coverprofile=../$(PRF_DIR)/coverage.out -cpuprofile=../$(PRF_DIR)/cpu.out -memprofile=../$(PRF_DIR)/mem.out
# $(GO_CMD) tool cover -html=coverage.out -o coverage.html
# $(GO_CMD) tool pprof cpu.out
# $(GO_CMD) tool pprof mem.out

# Clean build artifacts
clean:
	rm -f $(BIN_DIR)/$(BIN_NAME)*


.PHONY: all build clean
