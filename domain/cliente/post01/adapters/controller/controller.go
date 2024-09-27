package controller

import (
	"encoding/json"

	"github.com/valdinei-santos/api-modelo-clean-arch/config/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01/usecase"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// Execute - ...
func Execute(stamp string, c *gin.Context, useCase usecase.IUsecase) error {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "Controller - cliente/post01 - Execute"))

	var input *usecase.Request
	//var input interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		logger.Error("Erro Decode input", err, zap.String("id", stamp), zap.String("mtd", "Controller - cliente/post01 - Execute"))
		return err
	}
	logger.Info(input.CPF)

	err = useCase.Execute(stamp, input)
	if err != nil {
		logger.Error("Erro Execute useCase", err, zap.String("id", stamp), zap.String("mtd", "Controller - cliente/post01 - Execute"))
		return err
	}
	return nil
}
