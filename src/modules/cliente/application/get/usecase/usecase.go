package usecase

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	"go.uber.org/zap"
)

// UseCase - ...
type UseCase struct {
	Repo IRepository // aqui referencia a interface Repository do pg-repo dessa entity
}

func NewUseCase(r IRepository) *UseCase {
	return &UseCase{
		Repo: r,
	}
}

// Execute - ...
func (u *UseCase) Execute(stamp, cpf string) (*dto.Response, error) {
	c, err := u.Repo.FindById(stamp, cpf)
	if err != nil {
		logger.Error("Erro Cliente...", err, zap.String("id", stamp), zap.String("mtd", "cliente/get-cliente - UseCase - Execute"))
		return nil, err
	}

	tels, err := u.Repo.FindByIdTelefone(stamp, cpf)
	if err != nil {
		logger.Error("Erro Telefone...", err, zap.String("id", stamp), zap.String("mtd", "cliente/get-cliente - UseCase - Execute"))
		return nil, err
	}

	// Transforma da entidade entities.Cliente para o DTO Cliente
	var telefones []string
	for _, v := range tels {
		telefones = append(telefones, v.Numero)
	}
	cliente := &dto.Cliente{
		Nome:      c.Nome,
		DtNasc:    c.DtNasc,
		CPF:       c.Cpf,
		Telefones: telefones,
	}

	// Retorna o DTO Cliente com Telefones no DTO Response.
	result := &dto.Response{
		Cliente: *cliente,
		//Telefones: telefones,
	}
	return result, nil
}
