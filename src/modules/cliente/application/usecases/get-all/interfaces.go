package getall

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
)

// IUsecase - ...
type IUsecase interface {
	Execute(stamp string) (*dto.ResponseClientes, error)
}
