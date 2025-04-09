package controller

import (
	"net/http"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/get/usecase"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// ExecuteGet - ...
func ExecuteGet(stamp string, c *gin.Context, useCase usecase.IUsecase) error {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente/get01 - Controller - ExecuteGet01"))

	cpf := c.Param("cpf")
	logger.Info("CPF: "+cpf, zap.String("id", stamp), zap.String("mtd", "cliente/get01 - Controller - ExecuteGet01"))
	resp, err := useCase.Execute(stamp, cpf)
	if err != nil {
		logger.Error("Erro...", err, zap.String("id", stamp), zap.String("mtd", "cliente/get01 - Controller - ExecuteGet01"))
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
