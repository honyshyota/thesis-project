.PHONY: build
build:
		go build -o sim simulator/main.go
		./sim

.PHONY: api
api:
		go build -o api cmd/main.go
		./api

.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build