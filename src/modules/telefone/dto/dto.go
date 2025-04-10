package dto

type Telefone struct {
	//Id     int `json:"id"`
	CPF    string `json:"cpf"`
	Numero string `json:"numero"`
}

type Request struct {
	CPF    string `json:"cpf"`
	Numero string `json:"numero"`
}

type Response struct {
	CPF    string `json:"cpf"`
	Numero string `json:"numero"`
}

type ResponseAll struct {
	Telefones []Telefone `json:"telefones"`
}

// OutputDefault - Struct com a resposta da API
type OutputDefault struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
