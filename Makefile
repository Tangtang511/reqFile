BIN=bin

build:
	go build -v -o $(BIN)/reqFile ./cmd

container:
	docker build -t reqFile:dev .