#!make

.DEFAULT_GOAL := test_all

test_all:
	go test -v -coverpkg=./... -coverprofile=profile.cov ./... && go tool cover -func profile.cov && rm -rf profile.cov

build.up:
	go build -o up -ldflags="-s -w" cmd/server/main.go

build.cli:
	go build -o upcli -ldflags="-s -w" cmd/cli/main.go

run.up:
	./up

run.cli:
	./upcli

build.docker:
	docker build -t up:latest .

run.docker:
	docker run -p 8080:8080 up:latest

buildAndRun.up: build.up run.up
buildAndRun.docker: build.docker run.docker
buildAndRun.cli: build.cli run.cli

	
