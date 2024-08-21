package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/namanag0502/go-todo-server/pkg/models"
	"github.com/namanag0502/go-todo-server/pkg/utils"
)

type userHandler struct {
	repo models.UserRepository
}

func NewUserHandler(repo models.UserRepository) *userHandler {
	return &userHandler{repo: repo}
}

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.Find(r.Context())
	if err != nil {
		utils.WriteError(w, err, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	utils.WriteJSONResponse(w, users, "Users fetched successfully", http.StatusOK)
}

func (h *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParamFromCtx(r.Context(), "id")
	if id == "" {
		utils.WriteError(w, nil, "Missing user ID", http.StatusBadRequest)
		return
	}

	user, err := h.repo.FindOne(r.Context(), id)
	if err != nil {
		utils.WriteError(w, err, "Failed to fetch user", http.StatusInternalServerError)
		return
	}
	if user == nil {
		utils.WriteError(w, nil, "User not found", http.StatusNotFound)
		return
	}

	utils.WriteJSONResponse(w, user, "User fetched successfully", http.StatusOK)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParamFromCtx(r.Context(), "id")
	if id == "" {
		utils.WriteError(w, nil, "Missing user ID", http.StatusBadRequest)
		return
	}
	var req models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, err, "Failed to decode user", http.StatusBadRequest)
		return
	}

	result, err := h.repo.UpdateOne(r.Context(), id, req)
	if err != nil {
		utils.WriteError(w, err, "Failed to update user", http.StatusInternalServerError)
		return
	}

	if result == 0 {
		utils.WriteError(w, nil, "User not found", http.StatusNotFound)
		return
	}
	utils.WriteJSONResponse(w, nil, "User updated successfully", http.StatusOK)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParamFromCtx(r.Context(), "id")
	if id == "" {
		utils.WriteError(w, nil, "Missing user ID", http.StatusBadRequest)
		return
	}
	result, err := h.repo.DeleteOne(r.Context(), id)
	if err != nil {
		utils.WriteError(w, err, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	if result == 0 {
		utils.WriteError(w, nil, "User not found", http.StatusNotFound)
		return
	}

	utils.WriteJSONResponse(w, nil, "User deleted successfully", http.StatusOK)
}
