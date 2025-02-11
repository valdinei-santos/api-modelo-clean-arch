package usecase

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/domain/entities"
)

// IUsecase - ...
type IUsecase interface {
	Execute(stamp, cpf string) (*Response, error)
}

// IPresenter - Output Port
/* type IPresenter interface {
	Show(stamp string, t *Response) error
	ShowError(stamp string, msgErro string) error
} */

type IRepository interface {
	QueryLoadDataCliente(stamp, cpf string) (*entities.Cliente, error)
	QueryLoadDataTelefone(stamp, cpf string) ([]entities.Telefone, error)
}
