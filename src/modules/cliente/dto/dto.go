package dto

type Cliente struct {
	Nome   string `json:"nome"`
	DtNasc string `json:"dt_nasc"`
	CPF    string `json:"cpf"`
}

type ClienteComTelefone struct {
	Nome      string   `json:"nome"`
	DtNasc    string   `json:"dt_nasc"`
	CPF       string   `json:"cpf"`
	Telefones []string `json:"telefones"`
}

type Telefone struct {
	//Id     int `json:"id"`
	CPF    string `json:"cpf"`
	Numero string `json:"numero"`
}

type Request struct {
	CPF    string `json:"cpf"`
	Nome   string `json:"nome"`
	DtNasc string `json:"dt_nasc"`
}

type RequestComTelefone struct {
	CPF       string   `json:"cpf"`
	Nome      string   `json:"nome"`
	DtNasc    string   `json:"dt_nasc"`
	Telefones []string `json:"telefones"`
}

type Response struct {
	Cliente Cliente `json:"clientes"`
}

type ResponseComTelefone struct {
	CPF       string   `json:"cpf"`
	Nome      string   `json:"nome"`
	DtNasc    string   `json:"dt_nasc"`
	Telefones []string `json:"telefones"`
}

type ResponseClientes struct {
	Clientes []Cliente `json:"clientes"`
}

// OutputDefault - Struct com a resposta da API
type OutputDefault struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
