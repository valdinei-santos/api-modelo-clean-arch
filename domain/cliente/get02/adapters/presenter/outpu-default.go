package presenter

//OutputDefault - Struct com a resposta da API
type OutputDefault struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
