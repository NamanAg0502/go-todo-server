package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	return port
}

func WriteError(w http.ResponseWriter, err error, message string, statusCode int) {
	resp := &Response{
		Success: false,
		Message: message,
		Error:   err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if encodeErr := json.NewEncoder(w).Encode(resp); encodeErr != nil {
		http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
	}

}

func WriteJSONResponse(w http.ResponseWriter, data interface{}, message string, statusCode int) {
	resp := &Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if encodeErr := json.NewEncoder(w).Encode(resp); encodeErr != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
