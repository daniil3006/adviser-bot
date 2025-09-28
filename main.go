package main

import (
	tgClient "aviser-bot/clients/telegram"
	event_consumer "aviser-bot/consumer/event-consumer"
	"aviser-bot/events/telegram"
	"aviser-bot/storage/files"
	"flag"
	"log"
)

const (
	tgBOtHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBOtHost, mustToken()),
		files.New(storagePath),
	)

	log.Println("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
