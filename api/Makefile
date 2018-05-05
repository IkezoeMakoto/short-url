GLIDE=${GOPATH}/bin/glide

.PHONY: setup build run fmt

default: build

setup: .env $(GLIDE)
	$(GLIDE) install

build:
	go build -o bin/server src/main.go

run:
	go run src/main.go

fmt:
	find src/ -type f -name '*.go' | xargs -n 1 go fmt

vet:
	find src/ -type f -name '*.go' | xargs -n 1 go vet

.env:
	@cp .env.local .env

$(GLIDE):
	go get github.com/Masterminds/glide
	go install github.com/Masterminds/glide