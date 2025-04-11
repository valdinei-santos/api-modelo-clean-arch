package controller

import (
	"net/http"

	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/get-com-telefone"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// ExecuteGetComTelefone - ...
func ExecuteGetComTelefone(stamp string, c *gin.Context, useCase usecase.IUsecase) error {
	var dadosLog = "cliente - get-com-telefone - Controller"
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", dadosLog))

	cpf := c.Param("cpf")
	logger.Info("CPF: "+cpf, zap.String("id", stamp), zap.String("mtd", dadosLog))
	resp, err := useCase.Execute(stamp, cpf)
	if err != nil {
		logger.Error("Erro...", err, zap.String("id", stamp), zap.String("mtd", dadosLog))
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
