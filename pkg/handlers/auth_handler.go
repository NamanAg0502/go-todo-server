package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/namanag0502/go-todo-server/pkg/models"
	"github.com/namanag0502/go-todo-server/pkg/utils"
)

type authHandler struct {
	repo models.AuthRepository
}

func NewAuthHandler(repo models.AuthRepository) *authHandler {
	return &authHandler{repo}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, err, "Failed to decode login request", http.StatusBadRequest)
		return
	}

	user, err := h.repo.Login(r.Context(), &req)
	if err != nil {
		utils.WriteError(w, err, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		utils.WriteError(w, err, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	utils.WriteJSONResponse(w, token, "Logged in successfully", http.StatusOK)
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, err, "Failed to decode register request", http.StatusBadRequest)
		return
	}

	newUser, err := h.repo.Register(r.Context(), &req)
	if err != nil {
		utils.WriteError(w, err, "Failed to register user", http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateToken(newUser)
	if err != nil {
		utils.WriteError(w, err, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	utils.WriteJSONResponse(w, token, "User registered successfully", http.StatusCreated)
}
