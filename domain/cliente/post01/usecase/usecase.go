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
func (u *UseCase) Execute(stamp string, in *Request) error {
	log.Printf("%v - cliente/post01 - usecase - Execute", stamp)
	p := &Pessoa{
		CPF:    in.Pessoa.CPF,
		Nome:   in.Pessoa.Nome,
		DtNasc: in.Pessoa.DtNasc,
	}
	err := u.Repo.InsertPessoa(stamp, p)
	if err != nil {
		log.Printf("%v - ERRO-API: Erro em cliente/post01 Execute - %s", stamp, err.Error())
		u.Presenter.ShowError(stamp, "ERRO-API: Erro no Insert de dados!")
		return err
	}

	var result *Response
	result = &Response{
		StatusCode: 1,
		Message:    "Insert OK",
	}
	u.Presenter.Show(stamp, result)
	return nil
}
