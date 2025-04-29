package repository

import (
	"database/sql"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
)

type IRepository interface {
	BeginTransaction() (*sql.Tx, error)
	Save(p *dto.Cliente) error
	FindById(cpf string) (*entities.Cliente, error)
	FindAll() (*[]entities.Cliente, error)
	//FindByIdTelefone(stamp, cpf string) ([]entities.Telefone, error)
	//FindAllTelefone(stamp, cpf string) ([]entities.Telefone, error)
}
