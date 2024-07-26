package view

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01/adapters/presenter"

	"github.com/gin-gonic/gin"
)

type View struct {
	CtxGin *gin.Context
}

func NewView(c *gin.Context) *View {
	return &View{
		CtxGin: c,
	}
}

func (v *View) Show(stamp string, out *presenter.Output) error {
	log.Printf("%v - Show", stamp)
	FormatResponseToJSON(v.CtxGin.Writer, http.StatusOK, out)
	return nil
}

func (v *View) ShowError(stamp string, msgErro string) error {
	log.Printf("%v - ShowError", stamp)
	dataJErro := presenter.OutputDefault{
		StatusCode: -1,
		Message:    msgErro,
	}
	FormatResponseToJSON(v.CtxGin.Writer, http.StatusInternalServerError, dataJErro)
	return nil
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
