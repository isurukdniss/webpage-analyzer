# Build the binary
build:
	go build -o bin/web_analyzer main.go

# Run the application
run:
	go run main.go

# Run all tests
test:
	go test -cover ./...

# Clean the build directory
clean:
	rm -rf bin/