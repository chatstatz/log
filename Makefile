.PHONY: install build test

default:

install:
	@CGO_ENABLED=0 go get ./...

test:
	@CGO_ENABLED=0 go test -v ./...
