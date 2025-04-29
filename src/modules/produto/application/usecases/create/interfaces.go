package create

import "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"

// IUsecase - ...
type IUsecase interface {
	Execute(p *dto.Request) (*dto.OutputDefault, error)
}
