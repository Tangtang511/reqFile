BIN=bin

build:
	go build -v -o $(BIN)/reqFile .

container:
	docker build -t reqFile:dev .