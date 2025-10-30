package main

import (
	"log"
	"maxberries/catalog_service/internal/config"
	"maxberries/catalog_service/internal/infrastructure/postgres"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфига: %v", err)
	}
	log.Println("Конфиг успешно загружен")

	_, err = postgres.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к Postgres: %v", err)
	}
	log.Println("Мы успешно подключились к Postgres")

}
