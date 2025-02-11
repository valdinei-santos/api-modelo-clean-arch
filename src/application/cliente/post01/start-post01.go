package getativo

import (
	"database/sql"
	"time"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/application/cliente/post01/infra/controller"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/application/cliente/post01/infra/repository"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/application/cliente/post01/usecase"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func Start(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Start"))
	oraRepo := repository.NewRepoOracle(dbOra)
	u := usecase.NewUseCase(oraRepo)
	err := controller.Execute(stamp, ctx, u)
	if err != nil {
		logger.Error("Error", err, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Start"))
	}

}
