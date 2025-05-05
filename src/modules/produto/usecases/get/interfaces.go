package get

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
)

// IUsecase - ...
type IUsecase interface {
	Execute(id int) (*dto.Response, error)
}
