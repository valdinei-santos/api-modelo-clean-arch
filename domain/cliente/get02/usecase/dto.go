package usecase

/* type Request struct {
	Dado1      int     `json:"dado1"`
	Dado2      int     `json:"dado2"`
	Dado3      string  `json:"dado3"`
} */

/* type Telefone struct {
	//Id     int `json:"id"`
	Numero string `json:"numero"`
} */

type Cliente struct {
	Nome      string   `json:"nome"`
	DtNasc    string   `json:"dt_nasc"`
	CPF       string   `json:"cpf"`
	Telefones []string `json:"telefones"`
}

type Response struct {
	//Clientes *[]entities.ClienteComTel `json:"cliente"`
	Clientes []Cliente `json:"clientes"`
	//Telefones []string `json:"telefones"`
}

/* type ResponseAssisitdo struct {
	Dados        *DadosAssistido `json:"dados"`
	Movimentacao []*Movimentacao `json:"movimentacao"`
} */
