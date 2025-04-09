package usecase

import "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"

// IUsecase - ...
type IUsecase interface {
	Execute(stamp string, p *dto.Request) (*dto.OutputDefault, error)
}

type IRepository interface {
	InsertCliente(stamp string, p *dto.Cliente) error
	InsertTelefone(stamp string, t *dto.Telefone) error
}
