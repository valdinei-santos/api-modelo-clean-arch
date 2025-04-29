package repository

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
)

type IRepository interface {
	Save(t *dto.Telefone) error
	SaveAll(t []*dto.Telefone) error
	FindAll(cpf string) ([]entities.Telefone, error)
}
