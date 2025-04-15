package create

import (
	"log/slog"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/repository"
)

// UseCase - Estrutura para o caso de uso de criação de produto
type UseCase struct {
	Repo repository.IRepository // Interface do repositório para Produto
}

// NewUseCase - Construtor do caso de uso
func NewUseCase(r repository.IRepository) *UseCase {
	return &UseCase{
		Repo: r,
	}
}

// Execute - Executa a lógica de criação de um produto
func (u *UseCase) Execute(stamp string, in *dto.Request) (*dto.OutputDefault, error) {
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "produto - create - UseCase - Execute"))

	// Cria o objeto Produto a partir do DTO de entrada
	p := &dto.ProdutoDTO{
		Nome:       in.Nome,
		Descricao:  in.Descricao,
		Preco:      in.Preco,
		QtdEstoque: in.QtdEstoque,
		Categoria:  in.Categoria,
		FlAtivo:    in.FlAtivo,
	}

	// Salva o produto no repositório
	err := u.Repo.Save(stamp, p)
	if err != nil {
		slog.Error("Erro ao salvar produto", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "produto - create - UseCase - Execute"))
		return nil, err
	}

	// Retorna a resposta padrão
	result := &dto.OutputDefault{
		StatusCode: 1,
		Message:    "Produto inserido com sucesso",
	}
	return result, nil
}
