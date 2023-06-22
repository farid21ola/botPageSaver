package config

import (
	"flag"
	"log"
)

const (
	token      = ""
	connString = ""
)

type Config struct {
	TgToken                  string
	PostgresConnectionString string
}

func MustLoad() Config {
	token := flag.String("token-bot-token",
		token,
		"token for access to telegram bot",
	)
	flag.Parse()

	ConnectionString := flag.String(
		"postgresql-connection-string",
		connString,
		"connection string for Postgresql",
	)

	if *ConnectionString == "" {
		log.Fatal("connection string is not specified")
	}
	if *token == "" {
		log.Fatal("token is not specified")
	}

	return Config{
		*token,
		*ConnectionString,
	}
}
