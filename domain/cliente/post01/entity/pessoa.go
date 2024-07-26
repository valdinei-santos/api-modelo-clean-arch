package entity

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Pessoa struct {
	Cpf    string `json:"cpf"`
	Nome   string `json:"nome"`
	DtNasc string `json:"dt_nasc"`
}

func NewPessoa(
	cpf string,
	nome string,
	dtNasc string,
) (*Pessoa, error) {
	r2 := &Pessoa{
		Cpf:    cpf,
		Nome:   nome,
		DtNasc: dtNasc,
	}
	err := r2.Validate()
	if err != nil {
		return nil, errors.New("erro na construção do objeto Pessoa")
	}
	return r2, nil
}

// ValidatePessoa - ...
func (r *Pessoa) Validate() error {
	return validator.New().Struct(r)
}
