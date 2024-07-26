package controller

import (
	"log"

	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01/usecase"

	"github.com/gin-gonic/gin"
)

// Execute - ...
func Execute(stamp string, c *gin.Context, useCase usecase.IUsecase) error {
	log.Printf("%v - cliente/get01 - Controller - Execute", stamp)

	// Exemplo conversão String para Inteiro
	/* cpfInt, err := strconv.Atoi(c.Param("cpf"))
	if err != nil {
		return err
	} */

	cpf := c.Param("cpf")
	log.Printf("%v - cliente/get01 - Controller - Execute - CPF:%v", stamp, cpf)
	err := useCase.Execute(stamp, cpf)
	if err != nil {
		log.Printf("%v - Error: %v", stamp, err.Error())
		return err
	}
	return nil
}
