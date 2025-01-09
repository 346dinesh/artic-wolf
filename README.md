# Risk Management Service

This is a simple Golang service that provides basic Risk CRUD (Create and Read) operations.  
It listens on port 8080 for standard HTTP traffic.

## Requirements

- Go 1.18+ (or any recent Go version)
- No external database; data is stored in-memory.

## Installation

1. Clone this repository.
2. Ensure you have Go installed on your system.
3. Run `go mod tidy` to download required dependencies.

## Usage

1. Run the service:
   ```bash
   go run main.go
