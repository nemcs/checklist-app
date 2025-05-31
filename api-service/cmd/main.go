package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nemcs/checklist-app/api-service/internal/config"
	"github.com/nemcs/checklist-app/api-service/internal/handlers"
	"github.com/nemcs/checklist-app/api-service/internal/services"
	"log"
	"net/http"
)

// запуск сервера и роутинг api-service
func main() {
	//TODO вынести url в конфиг
	cfg := config.New()
	dbClient := services.NewDBClient(cfg.DBService.URL)
	h := handlers.NewHandler(dbClient)

	router := httprouter.New()
	//TODO подумать над названиями маршрутов
	router.GET("/list", h.ListTasks)
	router.POST("/tasks/:id", h.GetTask)
	router.PATCH("/done/:id", h.MarkDone)
	router.DELETE("/delete/:id", h.DeleteTask)
	router.POST("/create", h.CreateTask)

	log.Fatal(http.ListenAndServe(":9090", router))

}
