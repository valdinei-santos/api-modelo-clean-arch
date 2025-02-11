package controller

type Telefone struct {
	Numero string `json:"numero"`
}

type Telefones struct {
	Numeros []string `json:"numeros"`
}

type Pessoa struct {
	Nome      string   `json:"nome"`
	Cpf       string   `json:"cpf"`
	DtNasc    string   `json:"dt_nasc"`
	Telefones []string `json:"telefones"`
}

type Output struct {
	Pessoa *Pessoa `json:"pessoa"`
}

// OutputDefault - Struct com a resposta da API
type OutputDefault struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
