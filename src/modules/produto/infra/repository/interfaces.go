package repository

import (
	"database/sql"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
)

type IRepository interface {
	BeginTransaction() (*sql.Tx, error)
	Save(p *dto.ProdutoDTO) error
	FindById(id int) (*entities.Produto, error)
	FindAll() (*[]entities.Produto, error)
}
