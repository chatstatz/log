.SILENT:;
.PHONY: install test

install:
	CGO_ENABLED=0 go get ./...

test:
	CGO_ENABLED=1 go test -race -covermode=atomic ./...
