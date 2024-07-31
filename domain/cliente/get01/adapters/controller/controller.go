package controller

import (
	"log"

	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01/usecase"

	"github.com/gin-gonic/gin"
)

// Execute - ...
// @BasePath /api/modelo

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
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
