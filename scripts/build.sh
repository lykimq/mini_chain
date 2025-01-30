#! /bin/bash

# Exit on any error
set -e

# Build the application
echo "Building the application..."
go build -o mini-chain cmd/mini-chain/main.go

# Make the binary executable
chmod +x mini-chain
echo "Build completed successfully. Run ./mini-chain to start the application."

# Run the application
./mini-chain
echo "Application started successfully."

# Clean up
rm -f mini-chain
echo "Clean up completed."
