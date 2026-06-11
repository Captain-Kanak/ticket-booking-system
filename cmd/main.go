package main

import (
	"ticket-booking-system/internal/config"
	"ticket-booking-system/internal/server"
)

func main() {
	cfg := config.LoadEnv()

	db := config.ConnectDatabase(cfg)

	server.StartServer(cfg, db)
}
