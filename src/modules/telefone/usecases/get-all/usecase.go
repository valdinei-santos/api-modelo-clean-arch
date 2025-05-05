package getall

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"
)

// UseCase - ...
type UseCase struct {
	repo repository.IRepository // aqui referencia a interface Repository do pg-repo dessa entity
	log  logger.Logger
}

func NewUseCase(r repository.IRepository, l logger.Logger) *UseCase {
	return &UseCase{
		repo: r,
		log:  l,
	}
}

// Execute - ...
func (u *UseCase) Execute(cpf string) (*dto.ResponseAll, error) {
	u.log.Debug("Entrou getall.Execute")
	telefones, err := u.repo.FindAll(cpf)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "repo.FindAll")
		return nil, err
	}
	// Transforma a lista de entities.ClienteComTel para uma lista de DTO Cliente com respectivos telefones
	var tel dto.Telefone
	var tels []dto.Telefone
	for _, v := range telefones {
		tel = dto.Telefone{
			CPF:    v.Cpf,
			Numero: v.Numero,
		}
		tels = append(tels, tel)
	}
	result := &dto.ResponseAll{
		Telefones: tels,
	}
	return result, nil
}
