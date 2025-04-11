package get_com_telefone

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
)

// IUsecase - ...
type IUsecase interface {
	Execute(stamp, cpf string) (*dto.ResponseComTelefone, error)
}
