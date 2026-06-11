package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Port string
	Dsn  string
}

func LoadEnv() *EnvConfig {
	err := godotenv.Load()

	if err != nil {
		// panic(err)
		log.Fatal("Error: Loading .env file")
	}

	return &EnvConfig{
		Port: os.Getenv("PORT"),
		Dsn:  os.Getenv("DSN"),
	}
}
