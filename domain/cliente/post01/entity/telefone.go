package entity

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Telefone struct {
	Cpf    string `json:"cpf"`
	Numero string `json:"numero"`
}

func NewTelefone(
	id int,
	numero string,
	cpf string,
) (*Telefone, error) {
	r2 := &Telefone{
		Numero: numero,
		Cpf:    cpf,
	}
	err := r2.Validate()
	if err != nil {
		return nil, errors.New("erro na construção do objeto telefone")
	}
	return r2, nil
}

// ValidateTelefone - ...
func (r *Telefone) Validate() error {
	return validator.New().Struct(r)
}
