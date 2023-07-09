package account

import (
	"github.com/LulianoM/bank-api/src/controllers"
	"github.com/LulianoM/bank-api/src/http/dispatcher"
	"github.com/LulianoM/bank-api/src/interfaces"
)

type AccountHanlder struct {
	accountController interfaces.AccountController
}

func NewAccountHandler(controllers *controllers.ControllersContainer) *AccountHanlder {
	return &AccountHanlder{accountController: controllers.AccountController}
}

func (ah *AccountHanlder) Routes(d dispatcher.Dispatcher) {
	api := d.Public.Group("/account")
	api.Post("", ah.Create)
	api.Get("/:id", ah.Get)
}
