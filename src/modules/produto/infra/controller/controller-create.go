package controller

import (
	"encoding/json"
	"net/http"

	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/application/usecases/create"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// ExecuteCreate - Controlador para criar um produto
func ExecuteCreate(stamp string, c *gin.Context, useCase usecase.IUsecase) error {
	var dadosLog = "produto - create - Controller"
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", dadosLog))

	var input *dto.Request
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
