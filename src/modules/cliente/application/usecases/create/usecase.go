package create

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
func (u *UseCase) Execute(stamp string, in *dto.Request) (*dto.OutputDefault, error) {
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "cliente - create - UseCase - Execute"))
	p := &dto.Cliente{
		CPF:    in.CPF,
		Nome:   in.Nome,
		DtNasc: in.DtNasc,
	}
	err := u.Repo.Save(stamp, p)
	if err != nil {
		slog.Error("Erro - cliente - create - UseCase - Execute", slog.String("id", stamp), slog.Any("error", err))
		return nil, err
	}

	//var result *dto.Response
	result := &dto.OutputDefault{
		StatusCode: 1,
		Message:    "Insert OK",
	}
	//u.Presenter.Show(stamp, result)
	return result, nil
}
