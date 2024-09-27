package controller

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/config/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01/usecase"
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
	err := useCase.Execute(stamp, cpf)
	if err != nil {
		logger.Error("Erro...", err, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Controller - Execute"))
		return err
	}
	return nil
}
