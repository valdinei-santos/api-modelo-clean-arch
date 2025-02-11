package usecase

type Telefone struct {
	//Id     int `json:"id"`
	CPF    string `json:"cpf"`
	Numero string `json:"numero"`
}

type Cliente struct {
	CPF       string   `json:"cpf"`
	Nome      string   `json:"nome"`
	DtNasc    string   `json:"dt_nasc"`
	Telefones []string `json:"telefones"`
}

type Request struct {
	//Pessoa *Pessoa `json:"pessoa"`
	//Telefones []*string `json:"telefones"`
	CPF       string   `json:"cpf"`
	Nome      string   `json:"nome"`
	DtNasc    string   `json:"dt_nasc"`
	Telefones []string `json:"telefones"`
}

// Response - Output Data
type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
