package getativo

import (
	"database/sql"
	"time"

	"github.com/valdinei-santos/api-modelo-clean-arch/config/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01/adapters/controller"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01/adapters/presenter"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01/infra/repository"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01/infra/view"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01/usecase"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func Start(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Start"))
	oraRepo := repository.NewRepoOracle(dbOra)
	v := view.NewView(ctx)
	p := presenter.NewPresenter(v)
	u := usecase.NewUseCase(oraRepo, p)
	err := controller.Execute(stamp, ctx, u)
	if err != nil {
		logger.Error("Error", err, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Start"))
		p.ShowError(stamp, err.Error())
	}

}
