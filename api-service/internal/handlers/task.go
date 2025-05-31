package handlers

// HTTP-обработчики
import (
	"github.com/julienschmidt/httprouter"
	"github.com/nemcs/checklist-app/api-service/internal/constants"
	"github.com/nemcs/checklist-app/api-service/internal/models"
	"github.com/nemcs/checklist-app/api-service/internal/services"
	"github.com/nemcs/checklist-app/api-service/internal/utils"
	"log"
	"net/http"
)

type Handler struct {
	DB *services.DBClient
}

func NewHandler(db *services.DBClient) *Handler {
	return &Handler{DB: db}
}

// TODO возвращать в хендлерах инфу из db-service

func (h *Handler) ListTasks(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	tasks, err := h.DB.GetAllTasks()
	if err != nil {
		log.Printf("[Handler->ListTasks]: %v", err)
		utils.SendJSON(w, http.StatusInternalServerError, models.MessageResponse{Message: constants.ErrFetchingTasks})
		return
	}
	utils.SendJSON(w, http.StatusOK, tasks)
}

func (h *Handler) GetTask(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	task, err := h.DB.GetTask(params.ByName(constants.ID))
	if err != nil {
		log.Printf("[Handler->GetTask]: %v", err)
		utils.SendJSON(w, http.StatusInternalServerError, models.MessageResponse{Message: constants.ErrFetchingTask})
		return
	}
	utils.SendJSON(w, http.StatusOK, task)
}

func (h *Handler) MarkDone(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	if err := h.DB.ChangeStatusDone(params.ByName(constants.ID)); err != nil {
		log.Printf("[Handler->MarkDone]: %v", err)
		utils.SendJSON(w, http.StatusBadRequest, models.MessageResponse{Message: constants.ErrChangingStatus})
		return
	}
	utils.SendJSON(w, http.StatusOK, models.MessageResponse{Message: constants.MsgStatusChanged})
}

func (h *Handler) DeleteTask(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	if err := h.DB.DeleteTask(params.ByName(constants.ID)); err != nil {
		log.Printf("[Handler->DeleteTask]: %v", err)
		utils.SendJSON(w, http.StatusBadRequest, models.MessageResponse{Message: constants.ErrDeletingTask})
		return
	}
	utils.SendJSON(w, http.StatusOK, models.MessageResponse{Message: constants.MsgTaskDeleted})
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := h.DB.AddTask(r); err != nil {
		log.Printf("[Handler->CreateTask]: %v", err)
		utils.SendJSON(w, http.StatusBadRequest, models.MessageResponse{Message: constants.ErrCreatingTask})
		return
	}

	utils.SendJSON(w, http.StatusCreated, models.MessageResponse{Message: constants.MsgTaskCreated})
}
