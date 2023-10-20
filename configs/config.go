package configs

import (
	"effectiveMobile/pkg/util/log/sl"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type Config struct {
	DB                                   dbConf
	AgeHost, GenderHost, NationalityHost string
	ServerPort                           string
}

type dbConf struct {
	Username, Password string
	Host, Port, Name   string
}

func LoadConfigs(log *slog.Logger) *Config {
	if err := godotenv.Load(); err != nil {
		log.Error("could not load .env file", sl.Err(err))
		os.Exit(1)
	}

	return &Config{
		DB: dbConf{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		},
		AgeHost:         os.Getenv("AGE_HOST"),
		GenderHost:      os.Getenv("GENDER_HOST"),
		NationalityHost: os.Getenv("NATIONALITY_HOST"),

		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
