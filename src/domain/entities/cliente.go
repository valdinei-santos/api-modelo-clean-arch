package entities

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ClienteComTel struct {
	Cpf       string   `json:"cpf"`
	Nome      string   `json:"nome"`
	DtNasc    string   `json:"dt_nasc"`
	Telefones []string `json:"telefones"`
}

type Cliente struct {
	Cpf    string `json:"cpf"`
	Nome   string `json:"nome"`
	DtNasc string `json:"dt_nasc"`
}

func NewCliente(
	cpf string,
	nome string,
	dtNasc string,
) (*Cliente, error) {
	r2 := &Cliente{
		Cpf:    cpf,
		Nome:   nome,
		DtNasc: dtNasc,
	}
	err := r2.Validate()
	if err != nil {
		return nil, errors.New("erro na construção do objeto Cliente endpoint get02")
	}
	return r2, nil
}

// ValidateCliente - ...
func (r *Cliente) Validate() error {
	return validator.New().Struct(r)
}
