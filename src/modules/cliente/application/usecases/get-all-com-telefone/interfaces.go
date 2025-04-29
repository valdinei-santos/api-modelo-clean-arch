package getallcomtelefone

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
)

// IUsecase - ...
type IUsecase interface {
	Execute() ([]*dto.ResponseComTelefone, error)
}
