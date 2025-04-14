package getall

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/repository"
	"go.uber.org/zap"
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
	logger.Info("Entrou no caso de uso para obter todos os produtos", zap.String("id", stamp), zap.String("mtd", "produto/get-produtos - UseCase - Execute"))

	// Busca todos os produtos no repositório
	produtos, err := u.Repo.FindAll(stamp)
	if err != nil {
		logger.Error("Erro ao buscar produtos", err, zap.String("id", stamp), zap.String("mtd", "produto/get-produtos - UseCase - Execute"))
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
