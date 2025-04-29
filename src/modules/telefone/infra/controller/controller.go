package controller

import (
	"encoding/json"
	"net/http"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	create "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/application/usecases/create"
	getAll "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/application/usecases/get-all"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"

	"github.com/gin-gonic/gin"
)

// Create - ...
func Create(log logger.Logger, c *gin.Context, useCase create.IUsecase) {
	log.Debug("Entrou controller.Create")
	var input *dto.Request
	//var input interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		log.Error(err.Error(), "mtd", "json.NewDecoder")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, dataJErro)
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
		c.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
	return
}

// GetAll - ...
func GetAll(log logger.Logger, c *gin.Context, useCase getAll.IUsecase) {
	log.Debug("Entrou controller.GetAll")
	cpf := c.Param("cpf")
	resp, err := useCase.Execute(cpf)
	if err != nil {
		log.Error(err.Error(), "mtd", "useCase.Execute")
		dataJErro := dto.OutputDefault{
			StatusCode: -1,
			Message:    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, dataJErro)
		log.Info("### Finished ERROR", "status_code", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp)
	log.Info("### Finished OK", "status_code", http.StatusOK)
	return
}
