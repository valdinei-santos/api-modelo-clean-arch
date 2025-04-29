package cliente

import (
	"database/sql"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	create "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/create"
	createComTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/create-com-telefone"
	get "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/get"
	getAll "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/get-all"
	getAllComTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/get-all-com-telefone"
	getComTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/get-com-telefone"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/controller"
	repoCliente "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository"
	repoTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"

	"github.com/gin-gonic/gin"
)

func Create(log logger.Logger, ctx *gin.Context, dbOra *sql.DB) {
	log.Debug("Entrou cliente.Create")
	oraRepo := repoCliente.NewRepoOracle(dbOra, log)
	u := create.NewUseCase(oraRepo, log)
	controller.Create(log, ctx, u)
	/* if err != nil {
		log.Error(err.Error(), "mtd", "controller.Create", "status_code", http.StatusInternalServerError)
		return
	}
	log.Info("### Finished", "status_code", http.StatusOK) */
	return
}

func CreateComTelefone(log logger.Logger, ctx *gin.Context, dbOra *sql.DB) {
	log.Debug("Entrou cliente.CreateComTelefone")
	oraRepoCliente := repoCliente.NewRepoOracle(dbOra, log)
	oraRepoTelefone := repoTelefone.NewRepoOracle(dbOra, log)
	u := createComTelefone.NewUseCase(oraRepoCliente, oraRepoTelefone, log)
	controller.CreateComTelefone(log, ctx, u)
	return
}

func Get(log logger.Logger, c *gin.Context, dbOra *sql.DB) {
	log.Debug("Entrou cliente.Get")
	oraRepo := repoCliente.NewRepoOracle(dbOra, log)
	u := get.NewUseCase(oraRepo, log)
	controller.Get(log, c, u)
	return
}

func GetComTelefone(log logger.Logger, ctx *gin.Context, dbOra *sql.DB) {
	log.Debug("Entrou cliente.GetComTelefone")
	oraRepoCliente := repoCliente.NewRepoOracle(dbOra, log)
	oraRepoTelefone := repoTelefone.NewRepoOracle(dbOra, log)
	u := getComTelefone.NewUseCase(oraRepoCliente, oraRepoTelefone, log)
	controller.GetComTelefone(log, ctx, u)
	return
}

func GetAll(log logger.Logger, ctx *gin.Context, dbOra *sql.DB) {
	log.Debug("Entrou cliente.GetAll")
	oraRepo := repoCliente.NewRepoOracle(dbOra, log)
	u := getAll.NewUseCase(oraRepo, log)
	controller.GetAll(log, ctx, u)
	return
}

func GetAllComTelefone(log logger.Logger, ctx *gin.Context, dbOra *sql.DB) {
	log.Debug("Entrou cliente.GetAllComTelefone")
	oraRepoCliente := repoCliente.NewRepoOracle(dbOra, log)
	oraRepoTelefone := repoTelefone.NewRepoOracle(dbOra, log)
	u := getAllComTelefone.NewUseCase(oraRepoCliente, oraRepoTelefone, log)
	controller.GetAllComTelefone(log, ctx, u)
	return
}
