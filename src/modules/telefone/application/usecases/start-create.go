package usecases

import (
	"database/sql"
	"time"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/application/usecases/create"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/controller"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"

	//"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/create/usecase"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func StartCreate(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "start-create - telefone - StartCreate"))
	oraRepo := repository.NewRepoOracle(dbOra)
	u := usecase.NewUseCase(oraRepo)
	err := controller.ExecuteCreate(stamp, ctx, u)
	if err != nil {
		logger.Error("Error", err, zap.String("id", stamp), zap.String("mtd", "start-create - telefone - StartCreate"))
	}

}
