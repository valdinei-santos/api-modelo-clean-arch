package usecase

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/entities"
)

// // IUsecase - ...
// type IUsecase interface {
// 	Execute(stamp string) error
// }

// IPresenter - Output Port
type IPresenter interface {
	Show(stamp string, t *Response) error
	ShowError(stamp string, msgErro string) error
}

type IRepository interface {
	QueryLoadAllClientes(stamp string) (*[]entities.ClienteComTel, error)
	QueryLoadDataTelefone(stamp, cpf string) ([]entities.Telefone, error)
}
