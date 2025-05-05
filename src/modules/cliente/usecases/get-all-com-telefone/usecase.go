package getallcomtelefone

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository"
	repoTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"
)

// UseCase - ...
type UseCase struct {
	repoCliente  repository.IRepository   // aqui referencia a interface Repository desse recurso
	repoTelefone repoTelefone.IRepository // aqui referencia a interface Repository do recurso Telefone
	log          logger.Logger
}

func NewUseCase(repoCli repository.IRepository, repoTel repoTelefone.IRepository, l logger.Logger) *UseCase {
	return &UseCase{
		repoCliente:  repoCli,
		repoTelefone: repoTel,
		log:          l,
	}
}

// Execute - ...
func (u *UseCase) Execute() ([]*dto.ResponseComTelefone, error) {
	u.log.Debug("Entrou getallcomtelefone.Execute")
	clientes, err := u.repoCliente.FindAll()
	if err != nil {
		//u.log.Error("Error", err)
		u.log.Error(err.Error(), "mtd", "repoCliente.FindAll")
		return nil, err
	}

	var listaCli []*dto.ResponseComTelefone
	for _, v := range *clientes {
		tels, err := u.repoTelefone.FindAll(v.Cpf)
		if err != nil {
			//u.log.Error(err.Error())
			u.log.Error(err.Error(), "mtd", "repoTelefone.FindAll")
			return nil, err
		}
		telsStr := make([]string, len(tels))
		for i, t := range tels {
			telsStr[i] = t.Numero
		}
		oneCli := dto.ResponseComTelefone{
			Nome:      v.Nome,
			DtNasc:    v.DtNasc,
			CPF:       v.Cpf,
			Telefones: telsStr,
		}
		listaCli = append(listaCli, &oneCli)
	}

	return listaCli, nil
}
