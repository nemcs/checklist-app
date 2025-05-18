package main

// только запуск сервера и роутинг

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nemcs/checklist-app/api-service/internal/handlers"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.POST("/create", handlers.CreateTask)
	router.GET("/list", handlers.ListTasks)
	router.DELETE("/delete/:id", handlers.DeleteTask)
	router.PUT("/done/:id", handlers.MarkDone)
	log.Fatal(http.ListenAndServe(":9090", router))

}
