package controller

import (
	"net/http"
	"strconv"

	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/application/usecases/get"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// ExecuteGet - Controlador para obter um produto por ID
func ExecuteGet(stamp string, c *gin.Context, useCase usecase.IUsecase) error {
	var dadosLog = "produto - get - Controller"
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", dadosLog))

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Error("Erro ao converter ID", err, zap.String("id", stamp), zap.String("mtd", dadosLog))
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    "ID inválido",
		}
		c.JSON(http.StatusBadRequest, dataJErro)
		return err
	}

	logger.Info("ID: "+idParam, zap.String("id", stamp), zap.String("mtd", dadosLog))
	resp, err := useCase.Execute(stamp, id)
	if err != nil {
		logger.Error("Erro ao obter produto", err, zap.String("id", stamp), zap.String("mtd", dadosLog))
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
