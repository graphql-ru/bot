package main

import (
	"log"

	"github.com/graphql-ru/bot/handlers"
	"github.com/graphql-ru/bot/telegram"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Unable to load .env file")
	}

	client, err := telegram.New()

	if err != nil {
		panic("Telegram client failed")
	}

	client.Use(handlers.Logger)
	client.Use(handlers.Commands)

	client.Start()
}
