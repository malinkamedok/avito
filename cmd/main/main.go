package main

import (
	"avito/config"
	"avito/internal/app"
	"log"

	_ "avito/docs"
)

// @title Тестовое задание на позицию стажёра-бэкендера от Avito Tech
// @version 1.0
// @description Микросервис для работы с балансом пользователей

// @host localhost:9000
// @BasePath /
func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error in parse config: %s\n", err)
	}

	app.Run(cfg)
}
