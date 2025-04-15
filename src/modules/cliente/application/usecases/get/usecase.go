package get

import (
	"log/slog"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository"
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
func (u *UseCase) Execute(stamp, cpf string) (*dto.Response, error) {
	c, err := u.Repo.FindById(stamp, cpf)
	if err != nil {
		slog.Error("Erro Cliente...", err, slog.String("id", stamp), slog.String("mtd", "cliente/get-cliente - UseCase - Execute"))
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
