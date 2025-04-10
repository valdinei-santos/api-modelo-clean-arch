package create

import "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"

// IUsecase - ...
type IUsecase interface {
	Execute(stamp string, p *dto.Request) (*dto.OutputDefault, error)
}

/* type IRepository interface {
	Save(stamp string, p *dto.Telefone) error
} */
