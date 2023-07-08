package interfaces

import "github.com/LulianoM/bank-api/src/models"

type AccountController interface {
	Create(models.Account) error
}
