package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	create "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/create"
	createComTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/create-com-telefone"
	get "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/get"
	getAll "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/get-all"
	getAllComTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/get-all-com-telefone"
	getComTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/get-com-telefone"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"

	"github.com/gin-gonic/gin"
)

// ExecuteCreate - ...
func ExecuteCreate(stamp string, c *gin.Context, useCase create.IUsecase) error {
	var dadosLog = "cliente - create - Controller"
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", dadosLog))
	var input *dto.Request
	//var input interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		slog.Error("Erro Decode input", err, slog.String("id", stamp), slog.String("mtd", dadosLog))
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, dataJErro)
		return err
	}
	slog.Info(input.CPF)

	resp, err := useCase.Execute(stamp, input)
	if err != nil {
		slog.Error("Erro Execute useCase", err, slog.String("id", stamp), slog.String("mtd", dadosLog))
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

// ExecuteCreateComTelefone - ...
func ExecuteCreateComTelefone(stamp string, c *gin.Context, useCase createComTelefone.IUsecase) error {
	var dadosLog = "cliente - create-com-telefone - Controller"
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", dadosLog))
	var input *dto.RequestComTelefone
	//var input interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		slog.Error("Erro Decode input", err, slog.String("id", stamp), slog.String("mtd", dadosLog))
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, dataJErro)
		return err
	}

	resp, err := useCase.Execute(stamp, input)
	if err != nil {
		slog.Error("Erro Execute useCase", err, slog.String("id", stamp), slog.String("mtd", dadosLog))
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

// ExecuteGet - ...
func ExecuteGet(stamp string, c *gin.Context, useCase get.IUsecase) error {
	var dadosLog = "cliente - get - Controller"
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", dadosLog))
	cpf := c.Param("cpf")
	slog.Info("CPF: "+cpf, slog.String("id", stamp), slog.String("mtd", dadosLog))
	resp, err := useCase.Execute(stamp, cpf)
	if err != nil {
		slog.Error("Erro...", err, slog.String("id", stamp), slog.String("mtd", dadosLog))
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

// ExecuteGetComTelefone - ...
func ExecuteGetComTelefone(stamp string, c *gin.Context, useCase getComTelefone.IUsecase) error {
	var dadosLog = "cliente - get-com-telefone - Controller"
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", dadosLog))

	cpf := c.Param("cpf")
	slog.Info("CPF: "+cpf, slog.String("id", stamp), slog.String("mtd", dadosLog))
	resp, err := useCase.Execute(stamp, cpf)
	if err != nil {
		slog.Error("Erro...", err, slog.String("id", stamp), slog.String("mtd", dadosLog))
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
	var dadosLog = "cliente - get-all - Controller"
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", dadosLog))
	resp, err := useCase.Execute(stamp)
	if err != nil {
		slog.Error("Erro...", err, slog.String("id", stamp), slog.String("mtd", dadosLog))
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

// ExecuteGetComTelefone - ...
func ExecuteGetAllComTelefone(stamp string, c *gin.Context, useCase getAllComTelefone.IUsecase) error {
	var dadosLog = "cliente - get-all-com-telefone - Controller"
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", dadosLog))
	resp, err := useCase.Execute(stamp)
	if err != nil {
		slog.Error("Erro...", err, slog.String("id", stamp), slog.String("mtd", dadosLog))
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
