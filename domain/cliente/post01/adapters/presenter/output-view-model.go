package presenter

/* type Output struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
} */

type Telefone struct {
	Numero string `json:"numero"`
}

type Telefones struct {
	Numeros []Telefone `json:"numeros"`
}

type Pessoa struct {
	Nome      string    `json:"nome"`
	Cpf       string    `json:"cpf"`
	DtNasc    string    `json:"dt_nasc"`
	Telefones Telefones `json:"telefones"`
}

type Output struct {
	Pessoa *Pessoa `json:"pessoa"`
}
