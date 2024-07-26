package usecase

type Telefone struct {
	//Id     int `json:"id"`
	CPF    string `json:"cpf"`
	Numero string `json:"numero"`
}

type Pessoa struct {
	CPF    string `json:"cpf"`
	Nome   string `json:"nome"`
	DtNasc string `json:"dt_nasc"`
}

/* type Request struct {
	Pessoa    *Pessoa     `json:"pessoa"`
	Telefones []*Telefone `json:"telefones"`
} */

type Request struct {
	Pessoa    *Pessoa   `json:"pessoa"`
	Telefones []*string `json:"telefones"`
}

/* type Response struct {
	Pessoa    *Pessoa     `json:"pessoa"`
	Telefones []*Telefone `json:"telefones"`
} */

// Response - Output Data
type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
