package usecases

import (
	"database/sql"
	"time"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/get-all-com-telefone"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/controller"
	repoCliente "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository"
	repoTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func StartGetAllComTelefone(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente - get-all-com-telefone - StartGetAllComTelefone"))
	//oraRepo := repository.NewRepoOracle(dbOra)
	oraRepoCliente := repoCliente.NewRepoOracle(dbOra)
	oraRepoTelefone := repoTelefone.NewRepoOracle(dbOra)
	u := usecase.NewUseCase(oraRepoCliente, oraRepoTelefone)
	err := controller.ExecuteGetAllComTelefone(stamp, ctx, u)
	if err != nil {
		logger.Error("Error ", err, zap.String("id", stamp), zap.String("mtd", "cliente - get-all-com-telefone - StartGetAllComTelefone"))
		//p.ShowError(stamp, err.Error())
	}

}
