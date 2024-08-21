package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	DB *mongo.Database
}

func NewRouter(db *mongo.Database) *Router {
	return &Router{DB: db}
}

func (r *Router) InitRoutes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.StripSlashes)
	mux.Use(middleware.Recoverer)
	usersCollection := r.DB.Collection("users")
	todosCollection := r.DB.Collection("todos")

	mux.Route("/api/v1", func(r chi.Router) {
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(404)
			w.Write([]byte("Route does not exist"))
		})
		r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(405)
			w.Write([]byte("Method is not valid"))
		})
		r.Mount("/auth", authRoutes(usersCollection))
		r.Mount("/users", userRoutes(usersCollection))
		r.Mount("/todos", todoRoutes(todosCollection))
	})

	return mux
}
