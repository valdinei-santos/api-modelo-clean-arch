package dto

import "time"

// ProdutoDTO - Representa os dados básicos de um produto
type ProdutoDTO struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	QtdEstoque int     `json:"qtd_estoque"`
	Categoria  string  `json:"categoria"`
	FlAtivo    string  `json:"fl_ativo"`
}

// Request - Representa os dados necessários para criar ou atualizar um produto
type Request struct {
	Nome       string  `json:"nome" validate:"required"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco" validate:"required,gt=0"`
	QtdEstoque int     `json:"qtd_estoque" validate:"gte=0"`
	Categoria  string  `json:"categoria"`
	FlAtivo    string  `json:"fl_ativo"`
}

// Response - Representa a resposta de um produto único
type Response struct {
	ID              int       `json:"id"`
	Nome            string    `json:"nome"`
	Descricao       string    `json:"descricao"`
	Preco           float64   `json:"preco"`
	QtdEstoque      int       `json:"qtd_estoque"`
	Categoria       string    `json:"categoria"`
	DataCriacao     time.Time `json:"data_criacao"`
	DataAtualizacao time.Time `json:"data_atualizacao"`
	FlAtivo         string    `json:"fl_ativo"`
}

// ProdutosResponse - Representa a resposta de uma lista de produtos
type ProdutosResponse struct {
	Produtos []Response `json:"produtos"`
}

// OutputDefault - Struct com a resposta padrão da API
type OutputDefault struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
