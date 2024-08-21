package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/namanag0502/go-todo-server/pkg/handlers"
	"github.com/namanag0502/go-todo-server/pkg/middlewares"
	"github.com/namanag0502/go-todo-server/pkg/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func todoRoutes(r chi.Router, c *mongo.Collection) {
	s := services.NewTodoService(c)
	h := handlers.NewTodoHandler(s)

	r.Use(middlewares.AuthMiddleware)

	r.Get("/", h.GetAllTodos)
	r.Post("/", h.CreateTodo)
	r.Put("/{id}", h.UpdateTodoByID)
	r.Delete("/{id}", h.DeleteTodoByID)

}
