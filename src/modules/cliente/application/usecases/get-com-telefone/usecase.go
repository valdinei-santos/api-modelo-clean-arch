package get_com_telefone

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository"
	repoTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"
	"go.uber.org/zap"
)

// UseCase - ...
type UseCase struct {
	RepoCliente  repository.IRepository   // aqui referencia a interface Repository desse recurso
	RepoTelefone repoTelefone.IRepository // aqui referencia a interface Repository do recurso Telefone
}

func NewUseCase(repoCli repository.IRepository, repoTel repoTelefone.IRepository) *UseCase {
	return &UseCase{
		RepoCliente:  repoCli,
		RepoTelefone: repoTel,
	}
}

// Execute - ...
func (u *UseCase) Execute(stamp, cpf string) (*dto.ResponseComTelefone, error) {
	c, err := u.RepoCliente.FindById(stamp, cpf)
	if err != nil {
		logger.Error("Erro Cliente...", err, zap.String("id", stamp), zap.String("mtd", "cliente - UseCase - Execute"))
		return nil, err
	}

	tels, err := u.RepoTelefone.FindAll(stamp, cpf)
	if err != nil {
		logger.Error("Erro Telefone...", err, zap.String("id", stamp), zap.String("mtd", "cliente/get-cliente - UseCase - Execute"))
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
