package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Dsn  string
}

func LoadEnv() *Config {
	err := godotenv.Load()

	if err != nil {
		// panic(err)
		log.Fatal("Error: Loading .env file")
	}

	return &Config{
		Port: os.Getenv("PORT"),
		Dsn:  os.Getenv("DSN"),
	}
}
