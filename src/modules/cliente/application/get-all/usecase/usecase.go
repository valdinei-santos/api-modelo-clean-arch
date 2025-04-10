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
func (u *UseCase) Execute(stamp string) (*dto.ResponseClientes, error) {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente/get-clientes - UseCase - Execute"))
	clientes, err := u.Repo.FindAll(stamp)
	if err != nil {
		logger.Error("Erro...", err, zap.String("id", stamp), zap.String("mtd", "cliente/get-clientes - UseCase - Execute"))
		return nil, err
	}
	// Transforma a lista de entities.ClienteComTel para uma lista de DTO Cliente com respectivos telefones
	var cli dto.Cliente
	var listaCli []dto.Cliente
	for _, v := range *clientes {
		var tels []string
		tels = append(tels, v.Telefones...)
		cli = dto.Cliente{
			Nome:      v.Nome,
			DtNasc:    v.DtNasc,
			CPF:       v.Cpf,
			Telefones: tels,
		}
		listaCli = append(listaCli, cli)
	}
	result := &dto.ResponseClientes{
		Clientes: listaCli,
	}
	return result, nil
}
