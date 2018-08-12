install:
	dep ensure

build: install
	go build -o bin/bot

run:
	go run main.go

start: install build
	./bin/bot

dockerize: install
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/bot
	docker build -t tatyshev/graphql_ru_bot .
	docker push tatyshev/graphql_ru_bot
