package get

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
)

// IUsecase - ...
type IUsecase interface {
	Execute(stamp string, id int) (*dto.Response, error)
}
