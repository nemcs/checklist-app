package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/nemcs/checklist-app/db-service/internal/config"
	"github.com/nemcs/checklist-app/db-service/internal/models"
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

	dbpool, err := repository.NewPostgres(config.PostgresConfig{
		DBHost: conf.DBHost,
		DBPort: conf.DBPort,
		DBUser: conf.DBUser,
		DBPass: conf.DBPass,
		DBName: conf.DBName,
	})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
	defer dbpool.Close()
	fmt.Printf("DB connected: host=%s db=%s\n", conf.DBHost, conf.DBName)

	//Создание таблицы, если она еще не создана
	_, err = dbpool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS checklist (
	   id UUID PRIMARY KEY,
	   title TEXT NOT NULL,
	   description TEXT,
	   done BOOLEAN DEFAULT false,
	   created_at TIMESTAMP DEFAULT now()
	)`)
	if err != nil {
		log.Fatal("Ошибка при создании таблицы", err)
		return
	}

	repo := repository.NewTaskRepository(dbpool)
	//Добавляем новую таску
	task := models.Task{
		ID:          uuid.New().String(),
		Title:       "TEST19:35",
		Description: "TEST19:35",
	}
	err = repo.NewTask(context.Background(), task)
	if err != nil {
		log.Fatal("[NewTask] Ошибка при добавлении задачи:", err)
	}

	//Получаем одну таску по ID
	res, err := repo.GetTaskByID(context.Background(), "8a4deae0-19e7-4033-be37-fe51a4d55460")
	if err != nil {
		log.Fatal("[GetTaskByID] Ошибка при получении задачи:", err)
		return
	}
	fmt.Printf("ID: %v, Title: %s, Description: %s, Done: %v, CreatedAt: %v\n", res.ID, res.Title, res.Description, res.Done, res.CreatedAt)

	//Получаем все таски
	resAll, err := repo.GetAllTask(context.Background())
	if err != nil {
		log.Fatal("[GetAllTask] Ошибка при получении списка задач:", err)
		return
	}
	for _, task := range resAll {
		fmt.Printf("ID: %v, Title: %s, Description: %s, Done: %v, CreatedAt: %v\n", task.ID, task.Title, task.Description, task.Done, task.CreatedAt)
	}

	//Обновляем статус done на true
	var testIDforUpdate = "010c13b6-d13c-4d03-bf5c-e7543bf5cb1d"
	if err = repo.UpdateDoneByID(context.Background(), testIDforUpdate); err != nil {
		log.Fatal("[UpdateDoneByID] Ошибка при обновлении статуса Done:", err)
		return
	}
	//Получаем одну таску по ID
	res2, err := repo.GetTaskByID(context.Background(), testIDforUpdate)
	if err != nil {
		log.Fatal("[GetTaskByID] Ошибка при получении задачи:", err)
		return
	}
	fmt.Printf("ID: %v, Title: %s, Description: %s, Done: %v, CreatedAt: %v\n", res2.ID, res2.Title, res2.Description, res2.Done, res2.CreatedAt)

	//Удаляем задачу по айди
	if err = repo.DeleteTaskByID(context.Background(), "c5405ecc-bd13-44ae-b215-daca9b62ee84"); err != nil {
		log.Fatal("[DeleteTaskByID] Ошибка при удалении задачи", err)
		return
	}
}
