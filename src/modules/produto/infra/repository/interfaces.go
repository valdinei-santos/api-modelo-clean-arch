package repository

import (
	"database/sql"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
)

type IRepository interface {
	BeginTransaction(stamp string) (*sql.Tx, error)
	Save(stamp string, p *dto.ProdutoDTO) error
	FindById(stamp string, id int) (*entities.Produto, error)
	FindAll(stamp string) (*[]entities.Produto, error)
}
