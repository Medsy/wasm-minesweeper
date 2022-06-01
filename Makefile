
server: $(shell find ./server/ -type f)
	go build -o server ./server/main.go

wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm

build: server wasm

start: build
	server/main