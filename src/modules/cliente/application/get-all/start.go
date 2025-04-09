package usecases

import (
	"database/sql"
	"time"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/get-all/controller"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/get-all/repository"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/get-all/usecase"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func Start(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	oraRepo := repository.NewRepoOracle(dbOra)
	u := usecase.NewUseCase(oraRepo)
	err := controller.ExecuteGetAll(stamp, ctx, u)
	if err != nil {
		logger.Error("Erro...", err, zap.String("id", stamp), zap.String("mtd", "cliente02/get-clientes - Start"))
	}

}
