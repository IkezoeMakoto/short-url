.PHONY: build run fmt

default: build

fmt:
	find src/ -type f -name '*.go' | xargs -n 1 go fmt

build:
	go build -o bin/server src/main.go

run:
	go run src/main.go

.env:
	@cp .env.local .env