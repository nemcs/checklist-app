package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nemcs/checklist-app/db-service/internal/config"
	"github.com/nemcs/checklist-app/db-service/internal/repository"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	fmt.Println("db-services running")

	conf := config.New()

	conn, err := repository.NewPostgres(config.PostgresConfig{
		DBHost: conf.DBHost,
		DBPort: conf.DBPort,
		DBUser: conf.DBUser,
		DBPass: conf.DBPass,
		DBName: conf.DBName,
	})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
	defer conn.Close(context.Background())
	log.Printf("DB connected: host=%s db=%s", conf.DBHost, conf.DBName)

	conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS tasks (
	   id UUID PRIMARY KEY,
	   title TEXT NOT NULL,
	   description TEXT,
	   done BOOLEAN DEFAULT false,
	   created_at TIMESTAMP DEFAULT now()
	)`)
}
