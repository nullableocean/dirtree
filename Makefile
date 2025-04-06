.PHONY: build test

build:
	go build -o dirtree . 

test:
	go test -v