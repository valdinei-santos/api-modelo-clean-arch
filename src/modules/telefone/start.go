package telefone

import (
	"database/sql"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/controller"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"
	create "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/usecases/create"
	getAll "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/usecases/get-all"

	"github.com/gin-gonic/gin"
)

func Create(log logger.Logger, ctx *gin.Context, dbOra *sql.DB) {
	log.Debug("Entrou telefone.Create")
	oraRepo := repository.NewRepoOracle(dbOra, log)
	u := create.NewUseCase(oraRepo, log)
	controller.Create(log, ctx, u)
	return
}

func GetAll(log logger.Logger, ctx *gin.Context, dbOra *sql.DB) {
	log.Debug("Entrou telefone.GetAll")
	oraRepo := repository.NewRepoOracle(dbOra, log)
	u := getAll.NewUseCase(oraRepo, log)
	controller.GetAll(log, ctx, u)
	return

}
