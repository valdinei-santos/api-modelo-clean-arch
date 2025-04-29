package create

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/repository"
)

// UseCase - Estrutura para o caso de uso de criação de produto
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

// Execute - Executa a lógica de criação de um produto
func (u *UseCase) Execute(in *dto.Request) (*dto.OutputDefault, error) {
	u.log.Debug("Entrou create.Execute")

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
	err := u.repo.Save(p)
	if err != nil {
		u.log.Error(err.Error(), "mtd", "repo.Save")
		return nil, err
	}

	// Retorna a resposta padrão
	result := &dto.OutputDefault{
		StatusCode: 1,
		Message:    "Produto inserido com sucesso",
	}
	return result, nil
}
