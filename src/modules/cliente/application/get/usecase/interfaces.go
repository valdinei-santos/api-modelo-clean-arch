package usecase

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
)

// IUsecase - ...
type IUsecase interface {
	Execute(stamp, cpf string) (*dto.Response, error)
}

// IPresenter - Output Port
/* type IPresenter interface {
	Show(stamp string, t *Response) error
	ShowError(stamp string, msgErro string) error
} */

type IRepository interface {
	FindById(stamp, cpf string) (*entities.Cliente, error)
	FindByIdTelefone(stamp, cpf string) ([]entities.Telefone, error)
}
