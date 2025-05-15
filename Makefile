.PHONY: build run clean help fmt

# Binary name and directory
BINARY_NAME=famg
BIN_DIR=bin

# Default target
.DEFAULT_GOAL := help

# Help target
help:
	@echo "Available targets:"
	@echo "  make build     - Build the famg binary in $(BIN_DIR)/"
	@echo "  make run       - Build and run the famg binary"
	@echo "  make clean     - Remove the built binary"
	@echo "  make fmt       - Format Go code"
	@echo "  make help      - Show this help message"
	@echo "  make todo-test - Test creating a todo app"
# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BIN_DIR)
	@GO111MODULE=on go build -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/famg
	@echo "Build complete. Binary: $(BIN_DIR)/$(BINARY_NAME)"

# Run the application
run: build
	@echo "Running $(BINARY_NAME)..."
	@$(BIN_DIR)/$(BINARY_NAME)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR)
	@echo "Clean complete"

# Format Go code
fmt:
	@echo "Formatting Go code..."
	@go fmt ./...
	@echo "Format complete"

# Test the application
todo-test:
	@echo "E2E Testing Creating a todo app..."
	@rm -rf "../todo-app"
	@$(BIN_DIR)/$(BINARY_NAME) -path="../todo-app" -name="todo" -fullname="Todo App"
	@echo "Test complete"
