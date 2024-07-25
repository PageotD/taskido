# Variables
GO_CMD = go
SRC_DIR = sources
BIN_DIR = bin
BIN_NAME = taskido

# Targets
all: build

# Compile for Linux
build-linux:
	$(GO_CMD) -C $(SRC_DIR) build -o ../$(BIN_DIR)/$(BIN_NAME)-linux -ldflags="-s -w" -a -tags netgo -installsuffix netgo 

# Compile for Windows
build-windows:
	GOOS=windows GOARCH=amd64 $(GO_CMD) -C $(SRC_DIR) build -o ../$(BIN_DIR)/$(BIN_NAME).exe -ldflags="-s -w" -a -tags netgo -installsuffix netgo

# Build for both platforms
build: build-linux build-windows

# Clean build artifacts
clean:
	rm -f $(BIN_DIR)/$(BIN_NAME)-linux $(BIN_DIR)/$(BIN_NAME).exe


.PHONY: all build clean
