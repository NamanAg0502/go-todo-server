package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/namanag0502/go-todo-server/pkg/handlers"
	"github.com/namanag0502/go-todo-server/pkg/middlewares"
	"github.com/namanag0502/go-todo-server/pkg/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func userRoutes(r chi.Router, c *mongo.Collection) {
	s := services.NewUserService(c)
	h := handlers.NewUserHandler(s)

	r.Use(middlewares.AuthMiddleware)
	r.Get("/", h.GetUsers)
	r.Get("/{id}", h.GetUserByID)
	r.Put("/{id}", h.UpdateUser)
	r.Delete("/{id}", h.DeleteUser)
}
