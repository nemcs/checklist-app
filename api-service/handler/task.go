package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nemcs/checklist-app/api-service/service"
	"github.com/nemcs/checklist-app/api-service/util"
	"net/http"
)

//├── handler/                # HTTP-обработчики
//│   └── task.go             # все CRUD-функции по задачам

func CreateTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req, err := util.DecodeJsonBodyCreateTask(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := service.AddTask(*req)
	util.SendJSON(w, http.StatusCreated, task)
}
