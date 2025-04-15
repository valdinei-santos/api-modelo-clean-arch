package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	create "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/application/usecases/create"
	get "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/application/usecases/get"
	getAll "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/application/usecases/get-all"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
)

// ExecuteCreate - Controlador para criar um produto
func ExecuteCreate(stamp string, c *gin.Context, useCase create.IUsecase) error {
	var dadosLog = "produto - create - Controller"
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", dadosLog))
	var input *dto.Request
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		slog.Error("Erro Decode input", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", dadosLog))
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, dataJErro)
		return err
	}

	resp, err := useCase.Execute(stamp, input)
	if err != nil {
		slog.Error("Erro Execute useCase", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", dadosLog))
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

// ExecuteGet - Controlador para obter um produto por ID
func ExecuteGet(stamp string, c *gin.Context, useCase get.IUsecase) error {
	var dadosLog = "produto - get - Controller"
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", dadosLog))
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		slog.Error("Erro ao converter ID", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", dadosLog))
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    "ID inválido",
		}
		c.JSON(http.StatusBadRequest, dataJErro)
		return err
	}

	slog.Info("ID: "+idParam, slog.String("id", stamp), slog.String("mtd", dadosLog))
	resp, err := useCase.Execute(stamp, id)
	if err != nil {
		slog.Error("Erro ao obter produto", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", dadosLog))
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

// ExecuteGetAll - Controlador para obter todos os produtos
func ExecuteGetAll(stamp string, c *gin.Context, useCase getAll.IUsecase) error {
	var dadosLog = "produto - get-all - Controller"
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", dadosLog))

	resp, err := useCase.Execute(stamp)
	if err != nil {
		slog.Error("Erro ao obter produtos", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", dadosLog))
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
