package createcomtelefone

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository"
	dtoTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
	repoTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"
)

// UseCase - ...
type UseCase struct {
	repoCliente  repository.IRepository   // aqui referencia a interface Repository desse recurso
	repoTelefone repoTelefone.IRepository // aqui referencia a interface Repository do recurso Telefone
	log          logger.Logger
}

func NewUseCase(repoCli repository.IRepository, repoTel repoTelefone.IRepository, log logger.Logger) *UseCase {
	return &UseCase{
		repoCliente:  repoCli,
		repoTelefone: repoTel,
		log:          log,
	}
}

// Execute - ...
func (u *UseCase) Execute(in *dto.RequestComTelefone) (*dto.OutputDefault, error) {
	u.log.Debug("Entrou createcomtelefone.Execute")

	// Cria um DTO Cliente
	c := &dto.Cliente{
		CPF:    in.CPF,
		Nome:   in.Nome,
		DtNasc: in.DtNasc,
	}

	// Cria um slice de DTO Telefone
	var telefones []*dtoTelefone.Telefone
	for _, v := range in.Telefones {
		t := &dtoTelefone.Telefone{
			CPF:    in.CPF,
			Numero: v,
		}
		telefones = append(telefones, t)
	}

	// Inicia a transação para salvar o cliente e telefones na mesma transação
	tx, err := u.repoCliente.BeginTransaction()
	if err != nil {
		tx.Rollback()
		//u.log.Error(err.Error())
		u.log.Error(err.Error(), "mtd", "repoCliente.BeginTransaction")
		return nil, err
	}
	//defer tx.Rollback()

	// Salva o cliente
	err = u.repoCliente.Save(c)
	if err != nil {
		tx.Rollback()
		//u.log.Error(err.Error())
		u.log.Error(err.Error(), "mtd", "repoCliente.Save")
		return nil, err
	}

	// Salva os telefones
	err = u.repoTelefone.SaveAll(telefones)
	if err != nil {
		tx.Rollback()
		//u.log.Error(err.Error())
		u.log.Error(err.Error(), "mtd", "repoTelefone.SaveAll")
		return nil, err
	}

	// Commit da transação
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		//u.log.Error("Erro ao fazer commit: " + err.Error())
		u.log.Error(err.Error(), "mtd", "tx.Commit")
		return nil, err
	}

	// Monta o DTO de Response
	result := &dto.OutputDefault{
		StatusCode: 1,
		Message:    "Insert OK",
	}
	return result, nil
}
