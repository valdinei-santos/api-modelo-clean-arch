package usecases

import (
	"database/sql"
	"log/slog"
	"time"

	create "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/application/usecases/create"
	getAll "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/application/usecases/get-all"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/controller"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"

	"github.com/gin-gonic/gin"
)

func StartCreate(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "start-create - telefone - StartCreate"))
	oraRepo := repository.NewRepoOracle(dbOra)
	u := create.NewUseCase(oraRepo)
	err := controller.ExecuteCreate(stamp, ctx, u)
	if err != nil {
		slog.Error("Error", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "start-create - telefone - StartCreate"))
	}
}

func StartGetAll(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	slog.Info("start-get-all.go - StartGetAll")
	oraRepo := repository.NewRepoOracle(dbOra)
	u := getAll.NewUseCase(oraRepo)
	err := controller.ExecuteGetAll(stamp, ctx, u)
	if err != nil {
		slog.Error("Erro...", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "telefone/get-telefones - StartGetAll"))
	}

}
