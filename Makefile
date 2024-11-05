BINARY_NAME=haldi
BUILD_DIR=./build
DATA_DIR=/usr/local/haldi
BIN_DIR=/usr/local/bin
CONFIG_DIR=~/.haldi/
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

add-to-path:
	@if [ -z "$(FILE)" ]; then \
			echo "Error: file is not specified. Please provide the target file (e.g., FILE=~/.bash_profile)"; \
			exit 1; \
	fi
	@echo 'export PATH="$(DATA_DIR):$$PATH"' >> $(FILE)
	@echo 'Added your_directory_path to PATH in $(FILE)'

# TODO: check before creating if files/directories exist
install:
	@echo "Installing the binary to $(BIN_DIR)..."
	cp $(BUILD_DIR)/$(BINARY_NAME) $(BIN_DIR)/
	@echo "Creating haldi data directory $(DATA_DIR)..."
	sudo mkdir -p $(DATA_DIR)
	@echo "Creating haldi config file directory in home $(CONFIG_DIR)..."
	mkdir $(CONFIG_DIR)
	@echo "Creating haldi config file"
	cp $(CONFIG_FILE_SOURCE) $(CONFIG_DIR)/$(CONFIG_FILE)
	@echo "Installation complete! You can run '$(BINARY_NAME)' from the terminal."
	@echo "Haldi data directory '$(DATA_DIR)' was not added to your path and is required!"
	@echo "Please add it yourself or use command 'make add-to-path FILE=~/.zshrc'"

uninstall:
	@echo "Removing the binary from $(INSTALLATION_DESTINATION)..."
	rm -rf $(INSTALLATION_DESTINATION)
	sudo rm -rf $(DATA_DIR)
	rm -rf $(CONFIG_DIR)
	@echo "Uninstallation complete!"

default: build
