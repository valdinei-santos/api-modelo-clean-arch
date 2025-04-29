package create

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

func NewUseCase(r repository.IRepository, log logger.Logger) *UseCase {
	return &UseCase{
		repo: r,
		log:  log,
	}
}

// Execute - ...
func (u *UseCase) Execute(in *dto.Request) (*dto.OutputDefault, error) {
	u.log.Debug("Entrou create.Execute")
	t := &dto.Telefone{
		CPF:    in.CPF,
		Numero: in.Numero,
	}
	err := u.repo.Save(t)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "repo.Save")
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
