package create

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
func (u *UseCase) Execute(stamp string, in *dto.Request) (*dto.OutputDefault, error) {
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "telefone/create - UseCase - Execute"))
	//var telefones []string
	//for _, v := range in.Telefones {
	//telefones = append(telefones, in.Telefones...)
	//}
	t := &dto.Telefone{
		CPF:    in.CPF,
		Numero: in.Numero,
	}
	err := u.Repo.Save(stamp, t)
	if err != nil {
		slog.Error("Erro", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "telefone/create - UseCase - Execute"))
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
