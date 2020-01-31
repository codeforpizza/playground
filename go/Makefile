.PHONY: build
build:
	go build -v -o bin/playground cmd/playground/main.go

.PHONY: docker-build
docker-build:
	docker build . -t playground:golang

.PHONY: run
run:
	go run cmd/playground/main.go