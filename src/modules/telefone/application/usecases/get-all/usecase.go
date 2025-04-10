package getall

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"
	"go.uber.org/zap"
)

// UseCase - ...
type UseCase struct {
	Repo repository.IRepository // aqui referencia a interface Repository do pg-repo dessa entity
}

func NewUseCase(r repository.IRepository) *UseCase {
	return &UseCase{
		Repo: r,
	}
}

// Execute - ...
func (u *UseCase) Execute(stamp, cpf string) (*dto.ResponseAll, error) {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "Usecase - telefone - Execute"))
	telefones, err := u.Repo.FindAll(stamp, cpf)
	if err != nil {
		logger.Error("Erro...", err, zap.String("id", stamp), zap.String("mtd", "Usecase - telefone - Execute"))
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
