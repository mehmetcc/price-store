# Makefile

# Define constants
SRC = cmd/main.go
APP_NAME = symbol-store

# Default target
.PHONY: all
all: run

# Run the Go program
.PHONY: run
run:
	@echo "Running the Go program..."
	go run $(SRC)

# Build the binary
.PHONY: build
build:
	@echo "Building the binary..."
	go build -o $(APP_NAME) $(SRC)

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./...

# Clean up built files
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)
