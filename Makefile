install:
	dep ensure

build: install
	go build -o bin/bot
