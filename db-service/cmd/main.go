package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/nemcs/checklist-app/db-service/internal/config"
	"github.com/nemcs/checklist-app/db-service/internal/db"
	"github.com/nemcs/checklist-app/db-service/internal/handler"
	"github.com/nemcs/checklist-app/db-service/internal/repository"
	"log"
	"net/http"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	cfg := config.New()
	dbpool, err := db.NewPool(cfg.PostgresConfig)
	if err != nil {
		log.Fatal("Ошибка подключения к БД ", err)
	}
	log.Printf("БД подключена: host=%s db=%s\n", cfg.Host, cfg.DBName)
	defer dbpool.Close()

	log.Println("db-services running")

	repo := repository.New(dbpool)
	h := handler.New(repo)
	router := httprouter.New()
	router.GET("/tasks", h.GetAllTask)
	router.POST("/task", h.GetTaskByID)
	router.PATCH("/task/done", h.UpdateDoneByID)
	router.DELETE("/task/delete", h.DeleteTaskByID)
	router.POST("/task/create", h.NewTask)

	log.Fatal(http.ListenAndServe(":25432", router))

	//Создание таблицы, если она еще не создана
	// TODO вынести в миграции
	_, err = dbpool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS checklist (
	   id UUID PRIMARY KEY,
	   title TEXT NOT NULL,
	   description TEXT,
	   done BOOLEAN DEFAULT false,
	   created_at TIMESTAMP DEFAULT now()
	)`)
	if err != nil {
		log.Fatal("Ошибка при создании таблицы ", err)
		return
	}

}
