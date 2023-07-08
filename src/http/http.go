package http

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"github.com/LulianoM/bank-api/src/handlers"
	"github.com/LulianoM/bank-api/src/http/dispatcher"
)

type HTTPServer struct {
	Server *fiber.App
}

const (
	API_PATH   = "/api"
	ADMIN_PATH = "/admin"
)

func NewServer(handlers handlers.HandlerContainer) *HTTPServer {
	fiberServer := fiber.New()

	server := &HTTPServer{
		Server: fiberServer,
	}
	server.loadRoutes(handlers)
	return server
}

func (hs *HTTPServer) loadRoutes(handlers handlers.HandlerContainer) {

	publicRouter := hs.Server.Group(API_PATH)
	privateRouter := hs.Server.Group(ADMIN_PATH)
	routerDispatcher := dispatcher.Dispatcher{
		Public:  publicRouter,
		Private: privateRouter,
	}

	for _, handler := range handlers.Handlers {
		handler.Routes(routerDispatcher)
	}

}

func (hs *HTTPServer) Start() {
	server := hs.Server

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-quit
		fmt.Printf("Gracefully shutting down...")

		if err := server.Shutdown(); err != nil {
			fmt.Printf("Error while shutting down server: %s", err.Error())
		}
	}()

	if err := server.Listen(":8080"); err != nil {
		fmt.Printf("Error while starting up server: %s", err.Error())
	}

	fmt.Printf("Running cleanup tasks...")
}
