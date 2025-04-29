package getall

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository"
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
func (u *UseCase) Execute() (*dto.ResponseClientes, error) {
	u.log.Debug("Entrou getall.Execute")
	clientes, err := u.repo.FindAll()
	if err != nil {
		//u.log.Error(err.Error())
		u.log.Error(err.Error(), "mtd", "repo.FindAll")
		return nil, err
	}
	var cli dto.Cliente
	var listaCli []dto.Cliente
	for _, v := range *clientes {
		//var tels []string
		//tels = append(tels, v.Telefones...)
		cli = dto.Cliente{
			Nome:   v.Nome,
			DtNasc: v.DtNasc,
			CPF:    v.Cpf,
			//Telefones: tels,
		}
		listaCli = append(listaCli, cli)
	}
	result := &dto.ResponseClientes{
		Clientes: listaCli,
	}
	return result, nil
}
