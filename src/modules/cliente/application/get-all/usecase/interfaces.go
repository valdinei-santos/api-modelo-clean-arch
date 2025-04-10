package usecase

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
)

// IUsecase - ...
type IUsecase interface {
	Execute(stamp string) (*dto.ResponseClientes, error)
}

type IRepository interface {
	FindAll(stamp string) (*[]entities.ClienteComTel, error)
	FindAllTelefone(stamp, cpf string) ([]entities.Telefone, error)
}
