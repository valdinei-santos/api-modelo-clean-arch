package controller

import (
	"encoding/json"
	"log"

	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01/usecase"

	"github.com/gin-gonic/gin"
)

// Execute - ...
func Execute(stamp string, c *gin.Context, useCase usecase.IUsecase) error {
	log.Printf("%v - Controller - cliente/post01 - Execute", stamp)

	var input *usecase.Request
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		log.Printf("%v - Error: %v", stamp, err.Error())
		return err
	}

	err = useCase.Execute(stamp, input)
	if err != nil {
		log.Printf("%v - Error: %v", stamp, err.Error())
		return err
	}
	return nil
}
