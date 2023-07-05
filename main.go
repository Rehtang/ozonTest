package main

import (
	"flag"
	"log"

	"github.com/Rehtang/ozonTest/api"
	"github.com/Rehtang/ozonTest/storage"
)

func main() {
	// Флаги командной строки
	port := flag.String("port", "8080", "HTTP server port")
	storageType := flag.String("storage", "", "Storage type: inmemory or postgres")
	dbConnectionString := flag.String("postgres-link", "", "PostgreSQL database connection link")
	flag.Parse()

	// Создание хранилища
	var urlStorage storage.URLStorage
	switch *storageType {
	case "postgres":
		if *dbConnectionString == "" {
			log.Fatal("PostgreSQL connection string is required")
		}
		postgresStorage, err := storage.NewPostgresStorage(*dbConnectionString)
		if err != nil {
			log.Fatalf("Failed to initialize PostgreSQL storage: %v", err)
		}
		urlStorage = postgresStorage
	default:
		urlStorage = storage.NewInMemoryStorage()
	}

	// Создание контроллера и сервера
	handler := api.NewHandler(urlStorage)
	server := api.NewServer(handler)

	// Запуск сервера
	server.Start(*port)
}
