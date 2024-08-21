package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/namanag0502/go-todo-server/pkg/handlers"
	"github.com/namanag0502/go-todo-server/pkg/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func authRoutes(r chi.Router, c *mongo.Collection) {
	s := services.NewAuthService(c)
	h := handlers.NewAuthHandler(s)
	r.Post("/login", h.Login)
	r.Post("/register", h.Register)
}
