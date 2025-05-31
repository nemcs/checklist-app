package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nemcs/checklist-app/db-service/internal/models"
	"github.com/nemcs/checklist-app/db-service/internal/repository"
	"github.com/nemcs/checklist-app/db-service/internal/utils"
	"log"
	"net/http"
)

type Handler struct {
	repo *repository.ChecklistRepo
}

func New(repo *repository.ChecklistRepo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetAllTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tasks, err := h.repo.GetAllTask(r.Context())
	if err != nil {
		log.Printf("repo GetAllTask error: %v", err)
		utils.SendJSON(w, http.StatusBadRequest, models.MessageResponse{Message: "Не удалось получить задачи"})
		return
	}
	if len(tasks) == 0 {
		utils.SendJSON(w, http.StatusOK, models.MessageResponse{Message: "Список задач пуст"})
		return
	}
	utils.SendJSON(w, http.StatusOK, models.DataResponse{Data: tasks})
}

func (h *Handler) UpdateDoneByID(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id, err := utils.DecodeAndValidateIDFromBody(r)
	if err != nil {
		utils.SendJSON(w, http.StatusBadRequest, models.MessageResponse{Message: err.Error()})
		return
	}

	if err = h.repo.UpdateDoneByID(r.Context(), id); err != nil {
		log.Printf("repo UpdateDoneByID db error: %v", err)
		utils.SendJSON(w, http.StatusInternalServerError, models.MessageResponse{Message: "ошибка при обновлении статуса задачи"})
		return
	}

	//TODO сделать возврат всей таски
	utils.SendJSON(w, http.StatusOK, models.MessageResponse{Message: "Статус задачи изменён"})
}

func (h *Handler) DeleteTaskByID(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id, err := utils.DecodeAndValidateIDFromBody(r)
	if err != nil {
		utils.SendJSON(w, http.StatusBadRequest, models.MessageResponse{Message: err.Error()})
		return
	}
	if err = h.repo.DeleteTaskByID(r.Context(), id); err != nil {
		log.Printf("repo DeleteTaskByID db error: %v", err)
		utils.SendJSON(w, http.StatusInternalServerError, models.MessageResponse{Message: "ошибка при удалении задачи"})
		return
	}

	utils.SendJSON(w, http.StatusOK, models.MessageResponse{Message: "Задача удалена"})
}

func (h *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id, err := utils.DecodeAndValidateIDFromBody(r)
	if err != nil {
		utils.SendJSON(w, http.StatusBadRequest, models.MessageResponse{Message: err.Error()})
		return
	}

	task, err := h.repo.GetTaskByID(r.Context(), id)
	if err != nil {
		log.Printf("repo GetTaskByID db error: %v", err)
		utils.SendJSON(w, http.StatusInternalServerError, models.MessageResponse{Message: "ошибка при получении задачи по ID"})
		return
	}

	utils.SendJSON(w, http.StatusOK, models.DataResponse{Data: task})
}

// TODO
func (h *Handler) NewTask(_ http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		fmt.Errorf("Ошибка при создании задачи: %v", err)
		return
	}

	err := h.repo.NewTask(r.Context(), task)
	if err != nil {
		fmt.Errorf("Задача не создана: %v", err)
	}

}
