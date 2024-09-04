package main

import (
	"botFirst/Clients/telegram"
	event_consumer "botFirst/Consumer/event-consumer"
	telegram2 "botFirst/events/telegram"
	"botFirst/storage/Files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storageBath = "files_storage"
	batchSize   = 100
)

// Вот так
func main() {

	eventsProcessor := telegram2.New(telegram.New(tgBotHost, mustToken()), Files.NewStorage(storageBath))

	log.Print("started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service stopped", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	return *token
}
