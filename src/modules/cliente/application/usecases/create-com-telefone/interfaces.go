package createcomtelefone

import "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"

// IUsecase - ...
type IUsecase interface {
	Execute(p *dto.RequestComTelefone) (*dto.OutputDefault, error)
}
