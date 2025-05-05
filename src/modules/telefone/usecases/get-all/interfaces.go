package getall

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
)

// IUsecase - ...
type IUsecase interface {
	//Execute(stamp string) (*dto.Response, error)
	Execute(cpf string) (*dto.ResponseAll, error)
}

/* type IRepository interface {
	FindAll(stamp string) (*[]entities.ClienteComTel, error)
	FindAllTelefone(stamp, cpf string) ([]entities.Telefone, error)
} */
