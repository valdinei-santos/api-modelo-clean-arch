package controller

import (
	"encoding/json"
	"net/http"

	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/create"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// ExecuteCreate - ...
func ExecuteCreate(stamp string, c *gin.Context, useCase usecase.IUsecase) error {
	var dadosLog = "cliente - create - Controller"
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", dadosLog))

	var input *dto.Request
	//var input interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		logger.Error("Erro Decode input", err, zap.String("id", stamp), zap.String("mtd", dadosLog))
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, dataJErro)
		return err
	}
	logger.Info(input.CPF)

	resp, err := useCase.Execute(stamp, input)
	if err != nil {
		logger.Error("Erro Execute useCase", err, zap.String("id", stamp), zap.String("mtd", dadosLog))
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, dataJErro)
		return err
	}
	c.JSON(http.StatusOK, resp)
	return nil
}
