package repository

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
)

type IRepository interface {
	Save(stamp string, p *dto.Telefone) error
	SaveAll(stamp string, p []*dto.Telefone) error
	FindAll(stamp, cpf string) ([]entities.Telefone, error)
}
