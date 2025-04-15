package usecases

import (
	"database/sql"
	"log/slog"
	"time"

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

func StartGet(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "cliente.StartGet"))
	oraRepo := repoCliente.NewRepoOracle(dbOra)
	u := get.NewUseCase(oraRepo)
	err := controller.ExecuteGet(stamp, ctx, u)
	if err != nil {
		slog.Error("Error ", err, slog.String("id", stamp), slog.String("mtd", "cliente02.get-cliente.Start"))
	}
}

func StartGetComTelefone(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "cliente.StartGetComTelefone"))
	oraRepoCliente := repoCliente.NewRepoOracle(dbOra)
	oraRepoTelefone := repoTelefone.NewRepoOracle(dbOra)
	u := getComTelefone.NewUseCase(oraRepoCliente, oraRepoTelefone)
	err := controller.ExecuteGetComTelefone(stamp, ctx, u)
	if err != nil {
		slog.Error("Error ", err, slog.String("id", stamp), slog.String("mtd", "cliente02.get-cliente.Start"))
	}
}

func StartGetAll(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "cliente.StartGetAll"))
	oraRepo := repoCliente.NewRepoOracle(dbOra)
	u := getAll.NewUseCase(oraRepo)
	err := controller.ExecuteGetAll(stamp, ctx, u)
	if err != nil {
		slog.Error("Erro...", err, slog.String("id", stamp), slog.String("mtd", "cliente02/get-clientes - Start"))
	}
}

func StartGetAllComTelefone(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "cliente - get-all-com-telefone - StartGetAllComTelefone"))
	oraRepoCliente := repoCliente.NewRepoOracle(dbOra)
	oraRepoTelefone := repoTelefone.NewRepoOracle(dbOra)
	u := getAllComTelefone.NewUseCase(oraRepoCliente, oraRepoTelefone)
	err := controller.ExecuteGetAllComTelefone(stamp, ctx, u)
	if err != nil {
		slog.Error("Error ", err, slog.String("id", stamp), slog.String("mtd", "cliente - get-all-com-telefone - StartGetAllComTelefone"))
	}
}

func StartCreate(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "cliente - StartCreate"))
	oraRepo := repoCliente.NewRepoOracle(dbOra)
	u := create.NewUseCase(oraRepo)
	err := controller.ExecuteCreate(stamp, ctx, u)
	if err != nil {
		slog.Error("Error", err, slog.String("id", stamp), slog.String("mtd", "cliente02/create-cliente - Start"))
	}
}

func StartCreateComTelefone(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "cliente - StartCreateComTelefone"))
	oraRepoCliente := repoCliente.NewRepoOracle(dbOra)
	oraRepoTelefone := repoTelefone.NewRepoOracle(dbOra)
	u := createComTelefone.NewUseCase(oraRepoCliente, oraRepoTelefone)
	err := controller.ExecuteCreateComTelefone(stamp, ctx, u)
	if err != nil {
		slog.Error("Error", err, slog.String("id", stamp), slog.String("mtd", "cliente02/create-cliente - Start"))
	}
}
