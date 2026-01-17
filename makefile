.PHONY: run

run:
	go run cmd/main.go

build:
	go build -o backend-test cmd/main.go

clean:
	rm -rf backend-test
