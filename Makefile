#!make

.DEFAULT_GOAL := test_all

test_all:
	go test -v -coverpkg=./... -coverprofile=profile.cov ./... && go tool cover -func profile.cov && rm -rf profile.cov

build.server:
	go build -o server -ldflags="-s -w" cmd/server/main.go

build.cli:
	go build -o upcli -ldflags="-s -w" cmd/cli/main.go

run.server:
	./server

run.cli:
	./upcli

buildAndRun.server: build.server run.server
buildAndRun.cli: build.cli run.cli

	