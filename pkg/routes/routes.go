package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
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
		r.Route("/auth", func(r chi.Router) {
			authRoutes(r, usersCollection)
		})
		r.Route("/users", func(r chi.Router) {
			userRoutes(r, usersCollection)
		})
		r.Route("/todos", func(r chi.Router) {
			todoRoutes(r, todosCollection)
		})
	})

	return mux
}
