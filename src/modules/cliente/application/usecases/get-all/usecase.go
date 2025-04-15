package getall

import (
	"log/slog"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository"
)

// UseCase - ...
type UseCase struct {
	Repo repository.IRepository // aqui referencia a interface Repository do pg-repo dessa entity
}

func NewUseCase(r repository.IRepository) *UseCase {
	return &UseCase{
		Repo: r,
	}
}

// Execute - ...
func (u *UseCase) Execute(stamp string) (*dto.ResponseClientes, error) {
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "cliente/get-clientes - UseCase - Execute"))
	clientes, err := u.Repo.FindAll(stamp)
	if err != nil {
		slog.Error("Erro...", err, slog.String("id", stamp), slog.String("mtd", "cliente/get-clientes - UseCase - Execute"))
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
