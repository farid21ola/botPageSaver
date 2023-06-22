package main

import (
	tgClient "botPageSaver/clients/telegram"
	"botPageSaver/config"
	event_consumer "botPageSaver/consumer/event-consumer"
	"botPageSaver/events/telegram"
	storage "botPageSaver/storage"
	"botPageSaver/storage/postgres"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {
	cfg := config.MustLoad()

	var storage storage.Storage

	pool, err := postgres.NewPoolPostgres(cfg.PostgresConnectionString)
	if err != nil {
		log.Fatalln("ошибка инициализации БД: ", err)
	}
	storage = postgres.New(pool)

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, cfg.TgToken),
		storage,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
