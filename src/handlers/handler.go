package handlers

import (
	"github.com/LulianoM/bank-api/src/handlers/healthcheck"
	"github.com/LulianoM/bank-api/src/http/dispatcher"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Routes(dispatcher.Dispatcher)
}

type HandlerContainer struct {
	Handlers []Handler
	Server   *fiber.App
}

func NewHandlerContainer() HandlerContainer {
	return HandlerContainer{
		Handlers: []Handler{
			healthcheck.NewHealthHandler(),
		},
	}
}
