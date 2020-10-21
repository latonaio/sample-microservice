
DATE := $(shell date "+%Y%m%d%H%M")

all: build-arm64

build-arm64: build
	docker build -t sample-microservice-first:$(DATE) -f Dockerfile-first .
	docker tag sample-microservice-first:$(DATE) sample-microservice-first:latest
	docker build -t sample-microservice-second:$(DATE) -f Dockerfile-second .
	docker tag sample-microservice-second:$(DATE) sample-microservice-second:latest

build-amd64: build
	docker build -t sample-microservice-first:$(DATE) -f Dockerfile-first --build-arg ARCH=amd64 .
	docker tag sample-microservice-first:$(DATE) sample-microservice-first:latest
	docker build -t sample-microservice-second:$(DATE) -f Dockerfile-second --build-arg ARCH=amd64 .
	docker tag sample-microservice-second:$(DATE) sample-microservice-second:latest

build:
	go build ./cmd/first.go
	go build ./cmd/second.go

.PHONY: all
