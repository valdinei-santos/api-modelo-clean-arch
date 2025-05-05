package getall

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
)

// IUsecase - ...
type IUsecase interface {
	Execute() (*dto.ProdutosResponse, error)
}
