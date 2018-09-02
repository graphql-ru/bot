package main

import (
	"log"

	"github.com/graphql-ru/bot/handlers"
	"github.com/graphql-ru/bot/telegram"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	tg, err := telegram.New()

	if err != nil {
		log.Fatal("Telegram client failed")
	}

	tg.Use(handlers.Guard)
	tg.Use(handlers.Commands)
	tg.Use(handlers.Join)

	tg.Start()
}
