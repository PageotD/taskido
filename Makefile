# Variables
GO_CMD = go
SRC_DIR = sources
BIN_DIR = bin
BIN_NAME = taskido

# Targets
all: build

build:
	$(GO_CMD) build -C $(SRC_DIR) -o ../$(BIN_DIR)/$(BIN_NAME)

clean:
	rm -f $(BIN_DIR)/$(BIN_NAME)

.PHONY: all build clean
