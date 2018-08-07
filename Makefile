install:
	dep ensure

build: install
	go build -o bin/bot

run: 
	go run main.go

start: install build
	./bin/bot
