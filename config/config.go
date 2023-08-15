package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	databaseURLKey = "DATABASE_URL"
	tgtokenKey     = "TG_TOKEN"
)

type Config struct {
	TgToken                  string
	PostgresConnectionString string
}

func MustLoad() Config {
	if err := initConfig(); err != nil {
		log.Println("can't initialize config: ", err)
	}

	dbConfig := os.Getenv(databaseURLKey)
	if dbConfig == "" {
		log.Println("empty env config")
		dbConfig = viper.GetString(databaseURLKey)
	}

	tgToken := os.Getenv(tgtokenKey)
	if tgToken == "" {
		log.Println("empty tg token")
		tgToken = viper.GetString(tgtokenKey)
	}

	return Config{
		tgToken,
		dbConfig,
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("cfg")
	return viper.ReadInConfig()
}
