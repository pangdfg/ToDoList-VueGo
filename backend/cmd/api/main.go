package main

import (
	"log"
	"todolist-grpc/pkg/config"
	"todolist-grpc/pkg/db"

	"github.com/quangdangfit/gocommon/logger"
)

func main() {
	cfg := config.LoadConfig()
	_, err := db.Connect(cfg.DatabaseURI)
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}
	log.Println("Hello World")
}