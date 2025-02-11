package controller

import (
	"net/http"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/application/cliente/get01/usecase"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// Execute - ...
func Execute(stamp string, c *gin.Context, useCase usecase.IUsecase) error {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Controller - Execute"))

	// Exemplo conversão String para Inteiro
	/* cpfInt, err := strconv.Atoi(c.Param("cpf"))
	if err != nil {
		return err
	} */

	cpf := c.Param("cpf")
	logger.Info("CPF: "+cpf, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Controller - Execute"))
	resp, err := useCase.Execute(stamp, cpf)
	if err != nil {
		logger.Error("Erro...", err, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Controller - Execute"))
		dataJErro := OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, dataJErro)
		return err
	}
	c.JSON(http.StatusOK, resp)
	return nil
}
