package controller

import (
	"encoding/json"
	"net/http"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	create "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/usecases/create"
	createComTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/usecases/create-com-telefone"
	get "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/usecases/get"
	getAll "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/usecases/get-all"
	getAllComTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/usecases/get-all-com-telefone"
	getComTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/usecases/get-com-telefone"

	"github.com/gin-gonic/gin"
)

// Create - ...
func Create(log logger.Logger, ctx *gin.Context, useCase create.IUsecase) {
	log.Debug("Entrou controller.Create")
	var input *dto.Request
	err := json.NewDecoder(ctx.Request.Body).Decode(&input)
	if err != nil {
		log.Error(err.Error(), "mtd", "json.NewDecoder")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}
	log.Debug("CPF:" + input.CPF)

	resp, err := useCase.Execute(input)
	if err != nil {
		log.Error(err.Error(), "mtd", "useCase.Execute")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
	return
}

// CreateComTelefone - ...
func CreateComTelefone(log logger.Logger, ctx *gin.Context, useCase createComTelefone.IUsecase) {
	log.Debug("Entrou controller.CreateComTelefone")
	var input *dto.RequestComTelefone
	err := json.NewDecoder(ctx.Request.Body).Decode(&input)
	if err != nil {
		log.Error(err.Error(), "mtd", "json.NewDecoder")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}

	resp, err := useCase.Execute(input)
	if err != nil {
		log.Error(err.Error(), "mtd", "useCase.Execute")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
	return
}

// Get - ...
func Get(log logger.Logger, ctx *gin.Context, useCase get.IUsecase) {
	log.Debug("Entrou controller.Get")
	cpf := ctx.Param("cpf")
	log.Debug("CPF: " + cpf)
	resp, err := useCase.Execute(cpf)
	if err != nil {
		log.Error(err.Error(), "mtd", "useCase.Execute")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
	return
}

// GetComTelefone - ...
func GetComTelefone(log logger.Logger, ctx *gin.Context, useCase getComTelefone.IUsecase) {
	log.Debug("Entrou controller.GetComTelefone")
	cpf := ctx.Param("cpf")
	log.Debug("CPF: " + cpf)
	resp, err := useCase.Execute(cpf)
	if err != nil {
		log.Error(err.Error(), "mtd", "useCase.Execute")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
	return
}

// GetAll - ...
func GetAll(log logger.Logger, ctx *gin.Context, useCase getAll.IUsecase) {
	log.Debug("Entrou controller.GetAll")
	resp, err := useCase.Execute()
	if err != nil {
		log.Error(err.Error(), "mtd", "useCase.Execute")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
	return
}

// GetAllComTelefone - ...
func GetAllComTelefone(log logger.Logger, ctx *gin.Context, useCase getAllComTelefone.IUsecase) {
	log.Debug("Entrou controller.GetAllComTelefone")
	resp, err := useCase.Execute()
	if err != nil {
		log.Error(err.Error(), "mtd", "useCase.Execute")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
	return
}
