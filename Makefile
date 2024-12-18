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

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
	@echo "Clean up complete!"

.PHONY: install
install:
	@echo "Installing the binary to $(BIN_DIR)..."
	cp $(BUILD_DIR)/$(BINARY_NAME) $(BIN_DIR)/
	@echo "Creating haldi directory in $(CONFIG_DIR)..."
	mkdir -p $(CONFIG_DIR)/manifests
	@echo "Creating haldi config file"
	cp $(CONFIG_FILE_SOURCE) $(CONFIG_DIR)/$(CONFIG_FILE)
	@echo "Installation completed! You can run '$(BINARY_NAME)' from the terminal."

.PHONY: uninstall
uninstall:
	@echo "Removing haldi binary from $(INSTALLATION_DESTINATION)..."
	rm -rf $(INSTALLATION_DESTINATION)
	rm -rf $(CONFIG_DIR)
	@echo "Uninstallation completed!"

.PHONY: docs
docs:
	@echo "Runing docsify server..."
	docsify serve docs --port 6969

.PHONY: docs-install
docs-install:
	@echo "Installing docsify..."
	npm i docsify-cli -g

default: build
