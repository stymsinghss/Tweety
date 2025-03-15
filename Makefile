# Project settings
BINARY_NAME = tweety
PORT = 3000

# Build the Go binary
build:
	@echo "🔨 Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) ./main.go

# Run the project
run: build
	@echo "🚀 Starting $(BINARY_NAME) on port $(PORT)..."
	@./$(BINARY_NAME)

# Format Go code
fmt:
	@echo "🎨 Formatting code..."
	@go fmt ./...

# Clean up generated files
clean:
	@echo "🧹 Cleaning up..."
	@rm -f $(BINARY_NAME)

# Run everything (build, format, lint, and test)
all: fmt lint test build

# Help menu
help:
	@echo "Usage:"
	@echo "  make build   - Build the binary"
	@echo "  make run     - Run the project"
	@echo "  make fmt     - Format code"
	@echo "  make clean   - Remove generated files"