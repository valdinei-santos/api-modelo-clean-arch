package getcomtelefone

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
)

// IUsecase - ...
type IUsecase interface {
	Execute(cpf string) (*dto.ResponseComTelefone, error)
}
