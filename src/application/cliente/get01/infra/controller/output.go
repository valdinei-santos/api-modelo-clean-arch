package controller

import (
	"encoding/json"
	"net/http"
)

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

// OutputDefault - Struct com a resposta da API
type OutputDefault struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// FormatResponseToJSON - Para formatar a saida em JSON sem precisar criar uma Struct para isso
func FormatResponseToJSON(w http.ResponseWriter, statusCode int, response interface{}) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}
