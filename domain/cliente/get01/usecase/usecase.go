package usecase

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/config/logger"
	"go.uber.org/zap"
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
	c, err := u.Repo.QueryLoadDataCliente(stamp, cpf)
	if err != nil {
		logger.Error("Erro Cliente...", err, zap.String("id", stamp), zap.String("mtd", "cliente/get01 - UseCase - Execute"))
		u.Presenter.ShowError(stamp, "ERRO-API: Erro na carga de dados!")
		return err
	}

	tels, err := u.Repo.QueryLoadDataTelefone(stamp, cpf)
	if err != nil {
		logger.Error("Erro Telefone...", err, zap.String("id", stamp), zap.String("mtd", "cliente/get01 - UseCase - Execute"))
		u.Presenter.ShowError(stamp, "ERRO-API: Erro na carga de dados!")
		return err
	}

	// Transforma da entidade entities.Cliente para o DTO Cliente
	var telefones []string
	for _, v := range tels {
		telefones = append(telefones, v.Numero)
	}
	cliente := &Cliente{
		Nome:      c.Nome,
		DtNasc:    c.DtNasc,
		CPF:       c.Cpf,
		Telefones: telefones,
	}

	// Retorna o DTO Cliente com Telefones no DTO Response.
	result := &Response{
		Cliente: *cliente,
		//Telefones: telefones,
	}
	u.Presenter.Show(stamp, result)
	return nil
}
