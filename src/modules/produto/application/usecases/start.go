package usecases

import (
	"database/sql"
	"log/slog"
	"time"

	create "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/application/usecases/create"
	get "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/application/usecases/get"
	getAll "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/application/usecases/get-all"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/controller"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/repository"

	"github.com/gin-gonic/gin"
)

func StartCreate(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "produto - create - StartCreate"))
	oraRepo := repository.NewRepoOracle(dbOra)
	u := create.NewUseCase(oraRepo)
	err := controller.ExecuteCreate(stamp, ctx, u)
	if err != nil {
		slog.Error("Error", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "produto - create - Start"))
	}
}

func StartGet(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "produto.StartGet"))
	oraRepo := repository.NewRepoOracle(dbOra)
	u := get.NewUseCase(oraRepo)
	err := controller.ExecuteGet(stamp, ctx, u)
	if err != nil {
		slog.Error("Error ", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "produto.StartGet"))
		//p.ShowError(stamp, err.Error())
	}
}

func StartGetAll(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "produto.StartGetAll"))
	oraRepo := repository.NewRepoOracle(dbOra)
	u := getAll.NewUseCase(oraRepo)
	err := controller.ExecuteGetAll(stamp, ctx, u)
	if err != nil {
		slog.Error("Erro...", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "produto.StartGetAll"))
	}
}
