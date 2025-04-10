package usecases

import (
	"database/sql"
	"time"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	getall "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/application/usecases/get-all"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/controller"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func StartGetAll(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	logger.Info("start-get-all.go - StartGetAll")
	oraRepo := repository.NewRepoOracle(dbOra)
	u := getall.NewUseCase(oraRepo)
	err := controller.ExecuteGetAll(stamp, ctx, u)
	if err != nil {
		logger.Error("Erro...", err, zap.String("id", stamp), zap.String("mtd", "telefone/get-telefones - StartGetAll"))
	}

}
