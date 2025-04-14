package controller

import (
	"net/http"

	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/application/usecases/get-all"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// ExecuteGetAll - Controlador para obter todos os produtos
func ExecuteGetAll(stamp string, c *gin.Context, useCase usecase.IUsecase) error {
	var dadosLog = "produto - get-all - Controller"
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", dadosLog))

	resp, err := useCase.Execute(stamp)
	if err != nil {
		logger.Error("Erro ao obter produtos", err, zap.String("id", stamp), zap.String("mtd", dadosLog))
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
