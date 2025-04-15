package get

import (
	"log/slog"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/repository"
)

// UseCase - Estrutura para o caso de uso de obtenção de produto
type UseCase struct {
	Repo repository.IRepository // Interface do repositório para Produto
}

// NewUseCase - Construtor do caso de uso
func NewUseCase(r repository.IRepository) *UseCase {
	return &UseCase{
		Repo: r,
	}
}

// Execute - Executa a lógica para obter um produto por ID
func (u *UseCase) Execute(stamp string, id int) (*dto.Response, error) {
	p, err := u.Repo.FindById(stamp, id)
	if err != nil {
		slog.Error("Erro ao buscar produto", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "produto/get-produto - UseCase - Execute"))
		return nil, err
	}

	// Transforma da entidade Produto para o DTO ProdutoResponse
	produto := &dto.Response{
		ID:              p.ID,
		Nome:            p.Nome,
		Descricao:       p.Descricao,
		Preco:           p.Preco,
		QtdEstoque:      p.QtdEstoque,
		Categoria:       p.Categoria,
		DataCriacao:     p.DataCriacao,
		DataAtualizacao: p.DataAtualizacao,
		FlAtivo:         p.FlAtivo,
	}

	// Retorna o DTO ProdutoResponse
	return produto, nil
}
