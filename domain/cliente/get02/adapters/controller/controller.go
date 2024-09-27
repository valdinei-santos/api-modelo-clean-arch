package controller

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/config/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// Execute - ...
func Execute(stamp string, c *gin.Context, useCase IUsecase) error {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente/get02 - Controller - Execute"))
	err := useCase.Execute(stamp)
	if err != nil {
		logger.Error("Erro...", err, zap.String("id", stamp), zap.String("mtd", "cliente/get02 - Controller - Execute"))
		return err
	}
	return nil
}
