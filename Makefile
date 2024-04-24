.PHONY: test bench lint

test:
	go test -v ./...

bench:
	go test -bench=. -benchmem -v ./...

lint:
	golangci-lint run ./...

build:
	GOOS=linux GOARCH=amd64 go build -o app

