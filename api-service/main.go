package main

//├── main.go                 # только запуск сервера и роутинг

import (
	"github.com/nemcs/checklist-app/api-service/router"
	"log"
	"net/http"
)

func main() {
	r := router.Setup()
	log.Fatal(http.ListenAndServe(":8080", r))

}
