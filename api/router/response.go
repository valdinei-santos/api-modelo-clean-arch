package router

type ClienteResponse struct {
	Nome   string `json:"nome"`
	DtNasc string `json:"dt_nasc"`
	CPF    string `json:"cpf"`
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}
