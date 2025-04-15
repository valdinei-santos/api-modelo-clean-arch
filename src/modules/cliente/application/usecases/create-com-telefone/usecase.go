package create_com_telefone

import (
	"log/slog"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository"
	dtoTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
	repoTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"
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
func (u *UseCase) Execute(stamp string, in *dto.RequestComTelefone) (*dto.OutputDefault, error) {
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "cliente - create-com-telefone - UseCase - Execute"))

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
	tx, err := u.RepoCliente.BeginTransaction(stamp)
	if err != nil {
		slog.Error("Erro", err, slog.String("id", stamp), slog.String("mtd", "cliente - create-com-telefone - UseCase - BefginTransaction"))
		return nil, err
	}
	defer tx.Rollback()

	// Salva o cliente
	err = u.RepoCliente.Save(stamp, c)
	if err != nil {
		tx.Rollback()
		slog.Error("Erro", err, slog.String("id", stamp), slog.String("mtd", "cliente - create-com-telefone - UseCase - Execute"))
		return nil, err
	}

	// Salva os telefones
	err = u.RepoTelefone.SaveAll(stamp, telefones)
	if err != nil {
		tx.Rollback()
		slog.Error("Erro", err, slog.String("id", stamp), slog.String("mtd", "cliente - create-com-telefone - UseCase - Execute"))
		return nil, err
	}

	// Commit da transação
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		slog.Error("Erro ao fazer commit", err, slog.String("id", stamp), slog.String("mtd", "cliente - create-com-telefone - Usecase - Execute"))
		return nil, err
	}

	// Monta o DTO de Response
	result := &dto.OutputDefault{
		StatusCode: 1,
		Message:    "Insert OK",
	}
	return result, nil
}
