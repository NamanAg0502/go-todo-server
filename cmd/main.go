package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/namanag0502/go-todo-server/pkg/config"
	"github.com/namanag0502/go-todo-server/pkg/routes"
	"github.com/namanag0502/go-todo-server/pkg/utils"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	db := config.Init()

	mux := routes.NewRouter(db)

	port := utils.GetPort()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux.InitRoutes(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen to server on %s: %v\n", port, err)
		}
	}()
	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	config.Disconnect()

	log.Println("Server gracefully stopped")
}
