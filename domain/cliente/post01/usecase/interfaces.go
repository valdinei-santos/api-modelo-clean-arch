package usecase

// IUsecase - ...
type IUsecase interface {
	Execute(stamp string, p *Request) error
}

// IPresenter - Output Port
type IPresenter interface {
	Show(stamp string, t *Response) error
	ShowError(stamp string, msgErro string) error
}

type IRepository interface {
	InsertCliente(stamp string, p *Cliente) error
	InsertTelefone(stamp string, t *Telefone) error
}
