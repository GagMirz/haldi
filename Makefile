BINARY_NAME=haldi
BUILD_DIR=./build
BIN_DIR=/usr/local/bin
CONFIG_DIR=~/.haldi
CONFIG_FILE_SOURCE=./sample.config.json
CONFIG_FILE=config.json

.PHONY: build
build:
	@echo "Building the Go application..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME)
	@echo "Build complete! Binary is located at $(BUILD_DIR)/$(BINARY_NAME)"

clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
	@echo "Clean up complete!"

install:
	@echo "Installing the binary to $(BIN_DIR)..."
	cp $(BUILD_DIR)/$(BINARY_NAME) $(BIN_DIR)/
	@echo "Creating haldi directory in home $(CONFIG_DIR)..."
	mkdir -p $(CONFIG_DIR)/manifests
	@echo "Creating haldi config file"
	cp $(CONFIG_FILE_SOURCE) $(CONFIG_DIR)/$(CONFIG_FILE)
	@echo "Installation complete! You can run '$(BINARY_NAME)' from the terminal."

uninstall:
	@echo "Removing the binary from $(INSTALLATION_DESTINATION)..."
	rm -rf $(INSTALLATION_DESTINATION)
	rm -rf $(CONFIG_DIR)
	@echo "Uninstallation complete!"

default: build
