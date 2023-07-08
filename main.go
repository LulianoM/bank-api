package main

import (
	"github.com/LulianoM/bank-api/src/handlers"
	"github.com/LulianoM/bank-api/src/http"
)

func main() {
	handlersWithServices := handlers.NewHandlerContainer()
	http.NewServer(handlersWithServices).Start()
}
