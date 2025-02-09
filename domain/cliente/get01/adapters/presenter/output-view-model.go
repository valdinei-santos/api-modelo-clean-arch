package presenter

/* type Output struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
} */

type Telefone struct {
	Numero string `json:"numero"`
}

type Telefones struct {
	Numeros string `json:"numeros"`
}

type Cliente struct {
	Nome   string `json:"nome"`
	Cpf    string `json:"cpf"`
	DtNasc string `json:"dt_nasc"`
	//Telefones []Telefone `json:"telefones"`
	Telefones []string `json:"telefones"`
}

type Output struct {
	Cliente *Cliente `json:"cliente"`
}
