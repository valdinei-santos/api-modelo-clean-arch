package entities

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Telefone struct {
	Cpf    string
	Numero string
}

func NewTelefone(
	cpf string,
	numero string,
) (*Telefone, error) {
	r2 := &Telefone{
		Cpf:    cpf,
		Numero: numero,
	}
	err := r2.Validate()
	if err != nil {
		return nil, errors.New("erro na construção do objeto telefone do endpoint get02")
	}
	return r2, nil
}

// ValidateTelefone - ...
func (r *Telefone) Validate() error {
	return validator.New().Struct(r)
}
