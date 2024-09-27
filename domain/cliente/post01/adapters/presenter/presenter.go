package presenter

import (
	"encoding/json"

	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01/usecase"
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
	return out, nil
}
