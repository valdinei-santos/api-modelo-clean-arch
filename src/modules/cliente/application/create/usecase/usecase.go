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
func (u *UseCase) Execute(stamp string, in *dto.Request) (*dto.OutputDefault, error) {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente/post01 - UseCase - Execute"))
	var telefones []string
	//for _, v := range in.Telefones {
	telefones = append(telefones, in.Telefones...)
	//}
	p := &dto.Cliente{
		CPF:       in.CPF,
		Nome:      in.Nome,
		DtNasc:    in.DtNasc,
		Telefones: telefones,
	}
	err := u.Repo.InsertCliente(stamp, p)
	if err != nil {
		logger.Error("Erro", err, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - UseCase - Execute"))
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
