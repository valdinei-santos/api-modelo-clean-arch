package get_all_com_telefone

import (
	"log/slog"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository"
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
func (u *UseCase) Execute(stamp string) ([]*dto.ResponseComTelefone, error) {
	clientes, err := u.RepoCliente.FindAll(stamp)
	if err != nil {
		slog.Error("Erro Cliente...", err, slog.String("id", stamp), slog.String("mtd", "cliente - UseCase - Execute"))
		return nil, err
	}

	var listaCli []*dto.ResponseComTelefone
	for _, v := range *clientes {
		tels, err := u.RepoTelefone.FindAll(stamp, v.Cpf)
		if err != nil {
			slog.Error("Erro Telefone...", err, slog.String("id", stamp), slog.String("mtd", "cliente/get-cliente - UseCase - Execute"))
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
