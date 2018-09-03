package main

import (
	"log"
	"time"

	"github.com/graphql-ru/bot/gh"
	"github.com/graphql-ru/bot/handlers"
	"github.com/graphql-ru/bot/telegram"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	tgClient, err := telegram.New()
	ghClient := gh.New()

	if err != nil {
		log.Fatal("Telegram client failed")
	}

	tgClient.Use(handlers.Guard)
	tgClient.Use(handlers.Commands)
	tgClient.Use(handlers.Join)

	ghClient.ReminderTicker(10*time.Minute, func(msg string) {
		telegram.ToGraphQl(tgClient.Bot, msg)
	})

	tgClient.Start()
}
