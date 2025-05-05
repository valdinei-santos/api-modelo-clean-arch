package getcomtelefone

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
func (u *UseCase) Execute(cpf string) (*dto.ResponseComTelefone, error) {
	u.log.Debug("Entrou getcomtelefone.Execute")
	c, err := u.repoCliente.FindById(cpf)
	if err != nil {
		//u.log.Error(err.Error())
		u.log.Error(err.Error(), "mtd", "repoCliente.FindById")
		return nil, err
	}

	tels, err := u.repoTelefone.FindAll(cpf)
	if err != nil {
		//u.log.Error(err.Error())
		u.log.Error(err.Error(), "mtd", "repoTelefone.FindAll")
		return nil, err
	}

	// Transforma da entidade entities.Cliente para o DTO Cliente
	var telefones []string
	for _, v := range tels {
		telefones = append(telefones, v.Numero)
	}

	// Retorna o DTO Cliente com Telefones.
	result := &dto.ResponseComTelefone{
		Nome:      c.Nome,
		DtNasc:    c.DtNasc,
		CPF:       c.Cpf,
		Telefones: telefones,
	}
	return result, nil
}
