package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/devfullcycle/imersao22/go-gateway/internal/repository"
	"github.com/devfullcycle/imersao22/go-gateway/internal/service"
	"github.com/devfullcycle/imersao22/go-gateway/internal/web/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_SSL_MODE", "disable"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	accountRepository := repository.CreateAccountRepository(db)
	accountService := service.CreateAccountService(accountRepository)
	server := server.CreateServer(accountService, getEnv("PORT", "8080"))
	server.ConfigureRoutes()
	err = server.Start()
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
	fmt.Println("Server started on port", getEnv("PORT", "8080"))
	fmt.Println("Database connection established")
	fmt.Println("Environment variables loaded from .env file")
	fmt.Println("Server is running...")
	fmt.Println("Press Ctrl+C to stop the server")
	fmt.Println("Server is listening on port", getEnv("PORT", "8080"))
}
