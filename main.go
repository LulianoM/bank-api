package main

import (
	"github.com/LulianoM/bank-api/src/controllers"
	"github.com/LulianoM/bank-api/src/handlers"
	"github.com/LulianoM/bank-api/src/http"
)

func main() {
	controllersContainer := controllers.NewControllerContainer()
	handlersWithServices := handlers.NewHandlerContainer(controllersContainer)
	http.NewServer(handlersWithServices).Start()
}
