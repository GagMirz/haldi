BINARY_NAME=haldi
BUILD_DIR=./build
BIN_DIR=/usr/local/bin

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
	@echo "Installation completed! You can run '$(BINARY_NAME)' from the terminal."

.PHONY: uninstall
uninstall:
	@echo "Removing haldi binary from $(INSTALLATION_DESTINATION)..."
	rm -rf $(INSTALLATION_DESTINATION)
	@echo "Uninstallation completed!"

.PHONY: docs
docs:
	@echo "Runing docsify server..."
	docsify serve docs --port 6969

.PHONY: docs-install
docs-install:
	@echo "Installing docsify..."
	npm i docsify-cli -g

.PHONY: tag-and-release
tag-and-release:
	@if [ -z "$(TAG_NAME)" ]; then \
		echo "Error: You must specify a tag name."; \
		exit 1; \
	fi; \
	echo "Creating tag $(TAG_NAME)..."; \
	git tag $(TAG_NAME); \
	echo "Pushing tag $(TAG_NAME) to remote origin..."; \
	git push origin $(TAG_NAME)

default: build
