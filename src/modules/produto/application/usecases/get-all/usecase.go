package getall

import (
	"log/slog"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/repository"
)

// UseCase - Estrutura para o caso de uso de obtenção de todos os produtos
type UseCase struct {
	Repo repository.IRepository // Interface do repositório para Produto
}

// NewUseCase - Construtor do caso de uso
func NewUseCase(r repository.IRepository) *UseCase {
	return &UseCase{
		Repo: r,
	}
}

// Execute - Executa a lógica para obter todos os produtos
func (u *UseCase) Execute(stamp string) (*dto.ProdutosResponse, error) {
	slog.Info("Entrou no caso de uso para obter todos os produtos", slog.String("id", stamp), slog.String("mtd", "produto/get-produtos - UseCase - Execute"))

	// Busca todos os produtos no repositório
	produtos, err := u.Repo.FindAll(stamp)
	if err != nil {
		slog.Error("Erro ao buscar produtos", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "produto/get-produtos - UseCase - Execute"))
		return nil, err
	}

	// Transforma a lista de entidades Produto para uma lista de DTOs ProdutoResponse
	var listaProdutos []dto.Response
	for _, v := range *produtos {
		produto := dto.Response{
			ID:              v.ID,
			Nome:            v.Nome,
			Descricao:       v.Descricao,
			Preco:           v.Preco,
			QtdEstoque:      v.QtdEstoque,
			Categoria:       v.Categoria,
			DataCriacao:     v.DataCriacao,
			DataAtualizacao: v.DataAtualizacao,
			FlAtivo:         v.FlAtivo,
		}
		listaProdutos = append(listaProdutos, produto)
	}

	// Retorna a resposta com a lista de produtos
	result := &dto.ProdutosResponse{
		Produtos: listaProdutos,
	}
	return result, nil
}
