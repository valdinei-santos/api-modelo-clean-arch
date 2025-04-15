package getall

import (
	"log/slog"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"
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
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "Usecase - telefone - Execute"))
	telefones, err := u.Repo.FindAll(stamp, cpf)
	if err != nil {
		slog.Error("Erro...", err, slog.String("id", stamp), slog.String("mtd", "Usecase - telefone - Execute"))
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
