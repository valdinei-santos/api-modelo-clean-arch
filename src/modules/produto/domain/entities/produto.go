package entities

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type Produto struct {
	ID              int       `json:"id"` // ID gerado pelo banco de dados
	Nome            string    `json:"nome" validate:"required"`
	Descricao       string    `json:"descricao"`
	Preco           float64   `json:"preco" validate:"required,gt=0"`
	QtdEstoque      int       `json:"qtd_estoque" validate:"gte=0"`
	Categoria       string    `json:"categoria"`
	DataCriacao     time.Time `json:"data_criacao"`
	DataAtualizacao time.Time `json:"data_atualizacao"`
	FlAtivo         string    `json:"fl_ativo"`
}

func NewProduto(
	nome string,
	descricao string,
	preco float64,
	qtdEstoque int,
	categoria string,
	ativo string,
) (*Produto, error) {
	produto := &Produto{
		Nome:            nome,
		Descricao:       descricao,
		Preco:           preco,
		QtdEstoque:      qtdEstoque,
		Categoria:       categoria,
		DataCriacao:     time.Now(),
		DataAtualizacao: time.Now(),
		FlAtivo:         ativo,
	}
	err := produto.Validate()
	if err != nil {
		return nil, errors.New("erro na construção do objeto Produto")
	}
	return produto, nil
}

// Validate - Valida os campos do Produto
func (p *Produto) Validate() error {
	return validator.New().Struct(p)
}
