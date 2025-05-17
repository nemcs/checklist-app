package handlers

// HTTP-обработчики
import (
	"github.com/julienschmidt/httprouter"
	"github.com/nemcs/checklist-app/api-service/internal/services"
	"github.com/nemcs/checklist-app/api-service/internal/utils"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req, err := utils.DecodeJsonBodyCreateTask(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := services.AddTask(*req)
	utils.SendJSON(w, http.StatusCreated, task)
}
