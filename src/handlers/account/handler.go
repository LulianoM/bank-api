package account

import "github.com/LulianoM/bank-api/src/http/dispatcher"

type AccountHanlder struct {
}

func NewAccountHanlder() *AccountHanlder {
	return &AccountHanlder{}
}

func (ah *AccountHanlder) Routes(d dispatcher.Dispatcher) {
	api := d.Public.Group("/account")
	api.Post("", ah.Create)
}
