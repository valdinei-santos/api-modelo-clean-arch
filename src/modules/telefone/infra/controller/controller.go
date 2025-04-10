package controller

import (
	"encoding/json"
	"net/http"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/application/usecases/create"
	getAll "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/application/usecases/get-all"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// ExecuteCreate - ...
func ExecuteCreate(stamp string, c *gin.Context, useCase create.IUsecase) error {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "Controller - telefone/create - ExecuteCreate"))

	var input *dto.Request
	//var input interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		logger.Error("Erro Decode input", err, zap.String("id", stamp), zap.String("mtd", "Controller - telefone/create - ExecuteCreate"))
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, dataJErro)
		return err
	}

	resp, err := useCase.Execute(stamp, input)
	if err != nil {
		logger.Error("Erro Execute useCase", err, zap.String("id", stamp), zap.String("mtd", "Controller - telefone/create - ExecuteCreate"))
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

// ExecuteGetAll - ...
func ExecuteGetAll(stamp string, c *gin.Context, useCase getAll.IUsecase) error {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "Controller - telefone - ExecuteGetAll"))
	cpf := c.Param("cpf")
	/* var input *dto.Request
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		logger.Error("Erro Decode input", err, zap.String("id", stamp), zap.String("mtd", "Controller - telefone - ExecuteGetAll"))
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, dataJErro)
		return err
	} */

	resp, err := useCase.Execute(stamp, cpf)
	if err != nil {
		logger.Error("Erro...", err, zap.String("id", stamp), zap.String("mtd", "Controller - telefone - ExecuteGetAll"))
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
