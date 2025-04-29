package produto

import (
	"database/sql"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	create "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/application/usecases/create"
	get "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/application/usecases/get"
	getAll "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/application/usecases/get-all"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/controller"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/repository"

	"github.com/gin-gonic/gin"
)

func Create(log logger.Logger, ctx *gin.Context, dbOra *sql.DB) {
	log.Debug("Entrou produto.Create")
	oraRepo := repository.NewRepoOracle(dbOra, log)
	u := create.NewUseCase(oraRepo, log)
	controller.Create(log, ctx, u)
	return
}

func Get(log logger.Logger, ctx *gin.Context, dbOra *sql.DB) {
	log.Debug("Entrou produto.Get")
	oraRepo := repository.NewRepoOracle(dbOra, log)
	u := get.NewUseCase(oraRepo, log)
	controller.Get(log, ctx, u)
	return
}

func GetAll(log logger.Logger, ctx *gin.Context, dbOra *sql.DB) {
	log.Debug("Entrou usecases.GetAll")
	oraRepo := repository.NewRepoOracle(dbOra, log)
	u := getAll.NewUseCase(oraRepo, log)
	controller.GetAll(log, ctx, u)
	return
}
