package handlers

// HTTP-обработчики
import (
	"github.com/julienschmidt/httprouter"
	"github.com/nemcs/checklist-app/api-service/internal/models"
	"github.com/nemcs/checklist-app/api-service/internal/services"
	"github.com/nemcs/checklist-app/api-service/internal/utils"
	"net/http"
)

const (
	ID      = `id`
	MESSAGE = `message`
	CREATED = `created`
)

func CreateTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req, err := utils.DecodeJsonBodyCreateTask(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := services.AddTask(*req)
	utils.SendJSON(w, http.StatusCreated, map[string]string{
		ID:      task.ID,
		MESSAGE: CREATED,
	})
}

func ListTasks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tasks := services.GetAllTasks()
	if len(tasks) == 0 {
		utils.SendJSON(w, http.StatusOK, []models.Task{})
		return
	}
	utils.SendJSON(w, http.StatusOK, tasks)
}

func DeleteTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := services.DeleteTask(params.ByName(ID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func MarkDone(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := services.ChangeStatusDone(params.ByName(ID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
