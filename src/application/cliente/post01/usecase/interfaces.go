package usecase

// IUsecase - ...
type IUsecase interface {
	Execute(stamp string, p *Request) (*Response, error)
}

type IRepository interface {
	InsertCliente(stamp string, p *Cliente) error
	InsertTelefone(stamp string, t *Telefone) error
}
