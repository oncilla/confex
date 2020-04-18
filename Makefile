all: test lint build

lint:
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:v1.24.0 golangci-lint run -v

build:
	go build -o ./bin/confex -v

test:
	go test -v ./...

clean:
	go clean
	rm ./bin/*
