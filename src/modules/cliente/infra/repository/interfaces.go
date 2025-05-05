package repository

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
)

type IRepository interface {
	//BeginTransaction() (*sql.Tx, error)
	BeginTransaction() (ITransaction, error)
	Save(p *dto.Cliente) error
	FindById(cpf string) (*entities.Cliente, error)
	FindAll() (*[]entities.Cliente, error)
}

type ITransaction interface {
	Commit() error
	Rollback() error
}
