package presenter

import (
	"encoding/json"

	"github.com/valdinei-santos/api-modelo-clean-arch/config/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01/usecase"
	"go.uber.org/zap"
)

type Presenter struct {
	View IView
}

func NewPresenter(v IView) *Presenter {
	return &Presenter{
		View: v,
	}
}

func (p *Presenter) Show(stamp string, t *usecase.Response) error {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente/get01 - Presenter - Show"))
	saida, err := GetViewModelFromResponse(t)
	if err != nil {
		return err
	}
	err = p.View.Show(stamp, saida)
	if err != nil {
		return err
	}
	return nil
}

func (p *Presenter) ShowError(stamp string, msgErro string) error {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente/get01 - Presenter - ShowError"))
	err := p.View.ShowError(stamp, msgErro)
	if err != nil {
		return err
	}
	return nil
}

func GetViewModelFromResponse(res interface{}) (*Output, error) {
	var out *Output
	temporaryVariable, _ := json.Marshal(res)
	err := json.Unmarshal(temporaryVariable, &out)
	if err != nil {
		return nil, err
	}
	logger.Info("Response:"+out.Cliente.Nome, zap.String("id", ""), zap.String("mtd", "cliente/get01 - Presenter - GetViewModelFromResponse"))
	return out, nil
}
