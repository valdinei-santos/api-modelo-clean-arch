package get

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/repository"
)

// UseCase - Estrutura para o caso de uso de obtenção de produto
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

// Execute - Executa a lógica para obter um produto por ID
func (u *UseCase) Execute(id int) (*dto.Response, error) {
	u.log.Debug("Entrou get.Execute")
	p, err := u.repo.FindById(id)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "repo.FindById")
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
