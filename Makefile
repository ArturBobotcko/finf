# Define the name of the executable
APP_NAME = finf

# Define the target installation directory
INSTALL_DIR = /usr/local/bin

# Define the Go build command
GO_BUILD = go build

# Default target to build the application
.PHONY: all
all: build

# Build the executable
.PHONY: build
build:
	@echo "Building the application..."
	$(GO_BUILD) -o $(APP_NAME)

# Install the executable to /usr/local/bin
.PHONY: install
install: build
	@echo "Installing $(APP_NAME) to $(INSTALL_DIR)..."
	sudo cp $(APP_NAME) $(INSTALL_DIR)
	sudo chmod +x $(INSTALL_DIR)/$(APP_NAME)

# Remove the executable
.PHONY: clean
clean:
	@echo "Cleaning up the build..."
	rm -f $(APP_NAME)

# Run the application (useful for testing)
.PHONY: run
run: build
	@echo "Running the application..."
	./$(APP_NAME)

# Uninstall the executable from /usr/local/bin
.PHONY: uninstall
uninstall:
	@echo "Uninstalling $(APP_NAME) from $(INSTALL_DIR)..."
	sudo rm -f $(INSTALL_DIR)/$(APP_NAME)

