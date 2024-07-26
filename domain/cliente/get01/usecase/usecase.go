package usecase

import (
	"log"
)

// UseCase - ...
type UseCase struct {
	Repo      IRepository // aqui referencia a interface Repository do pg-repo dessa entity
	Presenter IPresenter
}

func NewUseCase(r IRepository, p IPresenter) *UseCase {
	return &UseCase{
		Repo:      r,
		Presenter: p,
	}
}

// Execute - ...
func (u *UseCase) Execute(stamp, cpf string) error {
	log.Printf("%v - cliente/get01 - usecase - Execute", stamp)
	c, err := u.Repo.QueryLoadDataCliente(stamp, cpf)
	if err != nil {
		log.Printf("%v - ERRO-API: Erro em cliente/get01 Execute - Cliente - %s", stamp, err.Error())
		u.Presenter.ShowError(stamp, "ERRO-API: Erro na carga de dados!")
		return err
	}
	cliente := &Cliente{
		Nome:   c.Nome,
		DtNasc: c.DtNasc,
		CPF:    c.Cpf,
	}

	tels, err := u.Repo.QueryLoadDataTelefone(stamp, cpf)
	if err != nil {
		log.Printf("%v - ERRO-API: Erro em cliente/get01 Execute - Telefone - %s", stamp, err.Error())
		u.Presenter.ShowError(stamp, "ERRO-API: Erro na carga de dados!")
		return err
	}

	var telefones []*Telefone
	for _, v := range tels {
		umDado := &Telefone{
			Numero: v.Numero,
		}
		log.Printf(umDado.Numero)
		telefones = append(telefones, umDado)
	}

	result := &Response{
		Cliente:   cliente,
		Telefones: telefones,
	}
	u.Presenter.Show(stamp, result)
	return nil
}
