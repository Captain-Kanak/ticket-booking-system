package main

import (
	"ticket-booking-system/internal/config"
	"ticket-booking-system/internal/server"
)

func main() {
	// * load environment variables
	cfg := config.LoadEnv()

	// * connect to database
	db := config.ConnectDatabase(cfg)

	// * start the server
	server.Start(cfg, db)
}
