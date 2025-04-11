package repository

import (
	"database/sql"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
)

type IRepository interface {
	BeginTransaction(stamp string) (*sql.Tx, error)
	Save(stamp string, p *dto.Cliente) error
	FindById(stamp, cpf string) (*entities.Cliente, error)
	FindAll(stamp string) (*[]entities.Cliente, error)
	//FindByIdTelefone(stamp, cpf string) ([]entities.Telefone, error)
	//FindAllTelefone(stamp, cpf string) ([]entities.Telefone, error)
}
