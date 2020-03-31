.DEFAULT_GOAL := run

prepare:
	go generate

run:
	go build -tags dev && ./corona

build-windows: prepare
	GOOS=windows GOARCH=amd64 go build -tags prod -ldflags "-w"

build-linux: prepare
	GOOS=linux GOARCH=amd64 go build -tags prod -ldflags "-w"

build-mac: prepare
	GOOS=darwin GOARCH=amd64 go build -tags prod -ldflags "-w"

