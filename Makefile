.PHONY: default usage build run build-release

usage:
	@echo "Please provide an option:"
	@echo "	make build		--- Build the app"
	@echo "	make run		--- Run the app"
	@echo "	make build-release	--- Build with optimization flags enabled"

build:
	go build -o server cmd/server/main.go

run:
	go run cmd/server/main.go

build-release:
	go build -o server -ldflags "-s -w" cmd/server/main.go

default: usage
