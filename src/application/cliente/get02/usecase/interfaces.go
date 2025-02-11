package usecase

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/domain/entities"
)

// IUsecase - ...
type IUsecase interface {
	Execute(stamp string) (*Response, error)
}

type IRepository interface {
	QueryLoadAllClientes(stamp string) (*[]entities.ClienteComTel, error)
	QueryLoadDataTelefone(stamp, cpf string) ([]entities.Telefone, error)
}
