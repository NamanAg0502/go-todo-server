package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/namanag0502/go-todo-server/pkg/handlers"
	"github.com/namanag0502/go-todo-server/pkg/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func authRoutes(c *mongo.Collection) *chi.Mux {
	r := chi.NewRouter()

	s := services.NewAuthService(c)
	h := handlers.NewAuthHandler(s)
	r.Post("/login", h.Login)
	r.Post("/register", h.Register)

	return r
}
