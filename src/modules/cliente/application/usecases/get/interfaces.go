package get

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
)

// IUsecase - ...
type IUsecase interface {
	Execute(stamp, cpf string) (*dto.Response, error)
}
