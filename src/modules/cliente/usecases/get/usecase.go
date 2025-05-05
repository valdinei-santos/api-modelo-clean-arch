package get

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository"
)

// UseCase - ...
type UseCase struct {
	repo repository.IRepository // aqui referencia a interface Repository do pg-repo dessa entity
	log  logger.Logger
}

func NewUseCase(r repository.IRepository, log logger.Logger) *UseCase {
	return &UseCase{
		repo: r,
		log:  log,
	}
}

// Execute - ...
func (u *UseCase) Execute(cpf string) (*dto.Response, error) {
	u.log.Debug("Entrou get.Execute")
	c, err := u.repo.FindById(cpf)
	if err != nil {
		//u.log.Error(err.Error())
		u.log.Error(err.Error(), "mtd", "repo.FindById")
		return nil, err
	}

	cliente := &dto.Cliente{
		Nome:   c.Nome,
		DtNasc: c.DtNasc,
		CPF:    c.Cpf,
	}

	// Retorna o DTO Cliente com Telefones no DTO Response.
	result := &dto.Response{
		Cliente: *cliente,
		//Telefones: telefones,
	}
	return result, nil
}
