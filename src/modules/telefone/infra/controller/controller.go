package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	create "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/application/usecases/create"
	getAll "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/application/usecases/get-all"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"

	"github.com/gin-gonic/gin"
)

// ExecuteCreate - ...
func ExecuteCreate(stamp string, c *gin.Context, useCase create.IUsecase) error {
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "Controller - telefone/create - ExecuteCreate"))
	var input *dto.Request
	//var input interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		slog.Error("Erro Decode input", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "Controller - telefone/create - ExecuteCreate"))
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, dataJErro)
		return err
	}

	resp, err := useCase.Execute(stamp, input)
	if err != nil {
		slog.Error("Erro Execute useCase", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "Controller - telefone/create - ExecuteCreate"))
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
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "Controller - telefone - ExecuteGetAll"))
	cpf := c.Param("cpf")

	resp, err := useCase.Execute(stamp, cpf)
	if err != nil {
		slog.Error("Erro...", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "Controller - telefone - ExecuteGetAll"))
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
