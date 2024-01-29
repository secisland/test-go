.PHONY: build

DEMO_VERSION=1.0.0
DEMO_ROOT=$(shell pwd)

build:
	GOOS=linux GOARCH=arm64 go build -o demo main.go
	docker build -f ./Dockerfile -t demo:$(DEMO_VERSION) $(DEMO_ROOT)