package usecase

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01/entity"
)

// IUsecase - ...
type IUsecase interface {
	Execute(stamp, cpf string) error
}

// IPresenter - Output Port
type IPresenter interface {
	Show(stamp string, t *Response) error
	ShowError(stamp string, msgErro string) error
}

type IRepository interface {
	QueryLoadDataCliente(stamp, cpf string) (*entity.Cliente, error)
	QueryLoadDataTelefone(stamp, cpf string) ([]*entity.Telefone, error)
}
