.PHONY: build
build:
		go build -o sim simulator/main.go
		./sim

.PHONY: api
api:
		go build -o api cmd/main.go
		./api

.DEFAULT_GOAL := build