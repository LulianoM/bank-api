package healthcheck

import (
	"github.com/LulianoM/bank-api/src/http/dispatcher"
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (this *HealthHandler) Routes(d dispatcher.Dispatcher) {
	api := d.Public.Group("/health")
	api.Get("/", this.Ping)
}

func (this *HealthHandler) Ping(c *fiber.Ctx) error {
	return c.Status(200).SendString("Pong :)")
}
