package usecases

import (
	"database/sql"
	"time"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/create-com-telefone"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/controller"
	repoCliente "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository"
	repoTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func StartCreateComTelefone(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente - StartCreateComTelefone"))
	oraRepoCliente := repoCliente.NewRepoOracle(dbOra)
	oraRepoTelefone := repoTelefone.NewRepoOracle(dbOra)
	u := usecase.NewUseCase(oraRepoCliente, oraRepoTelefone)
	err := controller.ExecuteCreateComTelefone(stamp, ctx, u)
	if err != nil {
		logger.Error("Error", err, zap.String("id", stamp), zap.String("mtd", "cliente02/create-cliente - Start"))
	}

}
