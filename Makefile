all: test build

build:
	go build -o ./bin/confex -v

test:
	go test -v ./...

clean:
	go clean
	rm ./bin/*
