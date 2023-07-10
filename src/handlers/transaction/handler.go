package transaction

import (
	"github.com/LulianoM/bank-api/src/controllers"
	"github.com/LulianoM/bank-api/src/http/dispatcher"
	"github.com/LulianoM/bank-api/src/interfaces"
)

type TransactionHanlder struct {
	transactionController interfaces.TransactionController
}

func NewTransactionHandler(controllers *controllers.ControllersContainer) *TransactionHanlder {
	return &TransactionHanlder{transactionController: controllers.TransactionController}
}

func (th *TransactionHanlder) Routes(d dispatcher.Dispatcher) {
	api := d.Public.Group("/transaction")
	api.Post("", th.Create)
}
