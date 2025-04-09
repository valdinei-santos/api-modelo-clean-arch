package usecases

import (
	"database/sql"
	"time"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/get/controller"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/get/repository"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/get/usecase"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func Start(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente02.get-cliente.Start"))
	oraRepo := repository.NewRepoOracle(dbOra)
	u := usecase.NewUseCase(oraRepo)
	err := controller.ExecuteGet(stamp, ctx, u)
	if err != nil {
		logger.Error("Error ", err, zap.String("id", stamp), zap.String("mtd", "cliente02.get-cliente.Start"))
		//p.ShowError(stamp, err.Error())
	}

}
