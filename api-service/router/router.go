package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nemcs/checklist-app/api-service/handler"
)

//└── router/                 # настройка роутера

func Setup() *httprouter.Router {
	router := httprouter.New()
	router.POST("/create", handler.CreateTask)
	// router.GET("/list", handler.ListTasks)
	// router.DELETE("/delete/:id", handler.DeleteTask)
	// router.PUT("/done/:id", handler.MarkDone)
	return router
}
