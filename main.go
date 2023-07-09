package main

import (
	"log"
	"os"

	"github.com/LulianoM/bank-api/internal/config"
	"github.com/LulianoM/bank-api/internal/database"
	"github.com/LulianoM/bank-api/src/controllers"
	"github.com/LulianoM/bank-api/src/handlers"
	"github.com/LulianoM/bank-api/src/http"
	"github.com/LulianoM/bank-api/src/repositories"
)

func main() {
	cfg := config.GetConfig()

	if err := cfg.Validate(); err != nil {
		log.Fatal(err, "invalid config in app initialization")
		os.Exit(1)
	}

	if err := database.Connect(cfg); err != nil {
		log.Fatal(err, "error to initialize database")
		os.Exit(1)
	}

	repositoriesContainer := repositories.NewRepositoriesContainer()
	controllersContainer := controllers.NewControllerContainer(repositoriesContainer)
	handlersWithServices := handlers.NewHandlerContainer(controllersContainer)
	http.NewServer(handlersWithServices).Start()
}
