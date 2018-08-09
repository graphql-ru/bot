package main

import (
	"log"

	"github.com/graphql-ru/bot/handlers"
	"github.com/graphql-ru/bot/telegram"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	client, err := telegram.New()

	if err != nil {
		log.Fatal("Telegram client failed")
	}

	client.Use(handlers.Logger)
	client.Use(handlers.Commands)

	client.Start()
}
