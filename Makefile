BINARY_NAME=haldi
BUILD_DIR=./build
DATA_DIR=/usr/local/haldi
BIN_LINK_DIR=/usr/local/bin

# VERSION=1.0.0

build:
	@echo "Building the Go application..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME)
	@echo "Build complete! Binary is located at $(BUILD_DIR)/$(BINARY_NAME)"

clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
	@echo "Clean up complete!"

add-to-path:
	@if [ -z "$(FILE)" ]; then \
			echo "Error: file is not specified. Please provide the target file (e.g., FILE=~/.bash_profile)"; \
			exit 1; \
	fi
	@echo 'export PATH="$(DATA_DIR):$$PATH"' >> $(FILE)
	@echo 'Added your_directory_path to PATH in $(FILE)'

install:
	@echo "Installing the binary to $(BIN_LINK_DIR)..."
	cp $(BUILD_DIR)/$(BINARY_NAME) $(BIN_LINK_DIR)/
	@echo "Creating haldi data directory $(DATA_DIR)..."
	sudo mkdir -p $(DATA_DIR)
	@echo "Installation complete! You can run '$(BINARY_NAME)' from the terminal."
	@echo "Haldi data directory '$(DATA_DIR)' was not added to your path and is required!"
	@echo "Please add it yourself or use command 'make add-to-path FILE=~/.zshrc'"

uninstall:
	@echo "Removing the binary from $(INSTALLATION_DESTINATION)..."
	rm -rf $(INSTALLATION_DESTINATION)
	sudo rmdir $(DATA_DIR)
	@echo "Uninstallation complete!"

default: build
