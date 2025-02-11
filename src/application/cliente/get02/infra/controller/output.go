package controller

type Cliente struct {
	Nome   string `json:"nome"`
	Cpf    string `json:"cpf"`
	DtNasc string `json:"dt_nasc"`
	//Telefones []Telefone `json:"telefones"`
	Telefones []string `json:"telefones"`
}

type Output struct {
	Clientes *[]Cliente `json:"clientes"`
}

// OutputDefault - Struct com a resposta da API
type OutputDefault struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
