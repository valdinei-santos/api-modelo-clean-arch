package usecase

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
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
func (u *UseCase) Execute(stamp string) (*Response, error) {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente/get02 - UseCase - Execute"))
	clientes, err := u.Repo.QueryLoadAllClientes(stamp)
	if err != nil {
		logger.Error("Erro...", err, zap.String("id", stamp), zap.String("mtd", "cliente/get02 - UseCase - Execute"))
		return nil, err
	}
	// Transforma a lista de entities.ClienteComTel para uma lista de DTO Cliente com respectivos telefones
	var cli Cliente
	var listaCli []Cliente
	for _, v := range *clientes {
		var tels []string
		tels = append(tels, v.Telefones...)
		cli = Cliente{
			Nome:      v.Nome,
			DtNasc:    v.DtNasc,
			CPF:       v.Cpf,
			Telefones: tels,
		}
		listaCli = append(listaCli, cli)
	}
	result := &Response{
		Clientes: listaCli,
	}
	return result, nil
}
