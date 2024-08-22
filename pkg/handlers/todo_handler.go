package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/namanag0502/go-todo-server/pkg/models"
	"github.com/namanag0502/go-todo-server/pkg/utils"
)

type TodoHandler struct {
	repo models.TodoRepository
}

func NewTodoHandler(repo models.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.repo.Find(r.Context())
	if err != nil {
		utils.WriteError(w, err, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}
	utils.WriteJSONResponse(w, todos, "Fetched todos successfully", http.StatusOK)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req models.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, err, "Failed to decode todo request", http.StatusBadRequest)
		return
	}
	todo, err := h.repo.CreateOne(r.Context(), req)
	if err != nil {
		utils.WriteError(w, err, "Failed to create todo", http.StatusInternalServerError)
		return
	}
	utils.WriteJSONResponse(w, todo, "Created todo successfully", http.StatusCreated)
}

func (h *TodoHandler) UpdateTodoByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParamFromCtx(r.Context(), "id")
	var req models.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, err, "Failed to decode todo request", http.StatusBadRequest)
		return
	}
	result, err := h.repo.UpdateOne(r.Context(), id, req)
	if err != nil {
		utils.WriteError(w, err, "Failed to update todo", http.StatusInternalServerError)
		return
	}
	if result == 0 {
		utils.WriteError(w, nil, "Todo not found", http.StatusNotFound)
		return
	}
	utils.WriteJSONResponse(w, nil, "Updated todo successfully", http.StatusOK)
}

func (h *TodoHandler) DeleteTodoByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParamFromCtx(r.Context(), "id")
	result, err := h.repo.DeleteOne(r.Context(), id)
	if err != nil {
		utils.WriteError(w, err, "Failed to delete todo", http.StatusInternalServerError)
		return
	}
	if result == 0 {
		utils.WriteError(w, nil, "Todo not found", http.StatusNotFound)
		return
	}
	utils.WriteJSONResponse(w, nil, "Deleted todo successfully", http.StatusOK)
}
