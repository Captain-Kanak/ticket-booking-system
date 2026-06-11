package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config *EnvConfig) *gorm.DB {
	dsn := config.Dsn

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connected successfully!")

	return db
}
