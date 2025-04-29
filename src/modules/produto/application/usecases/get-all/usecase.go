package getall

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/repository"
)

// UseCase - Estrutura para o caso de uso de obtenção de todos os produtos
type UseCase struct {
	repo repository.IRepository // Interface do repositório para Produto
	log  logger.Logger
}

// NewUseCase - Construtor do caso de uso
func NewUseCase(r repository.IRepository, l logger.Logger) *UseCase {
	return &UseCase{
		repo: r,
		log:  l,
	}
}

// Execute - Executa a lógica para obter todos os produtos
func (u *UseCase) Execute() (*dto.ProdutosResponse, error) {
	u.log.Debug("Entrou getall.Execute")

	// Busca todos os produtos no repositório
	produtos, err := u.repo.FindAll()
	if err != nil {
		u.log.Error(err.Error(), "mtd", "repo.FindAll")
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
