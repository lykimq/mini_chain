.PHONY: build test clean run

# Build the application
build:
	go build -o mini_chain ./cmd/mini_chain/

# Run tests
test:
	go test -v ./...

# Run test coverage
test-coverage:
	go test -v -cover ./...

# Clean up
clean:
	rm -f mini_chain

# Run the application
run: build
	./mini_chain

# Install dependencies
install:
	go mod tidy
