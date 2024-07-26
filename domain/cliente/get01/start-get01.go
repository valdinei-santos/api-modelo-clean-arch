package getativo

import (
	"database/sql"
	"log"
	"time"

	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01/adapters/controller"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01/adapters/presenter"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01/infra/repository"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01/infra/view"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01/usecase"

	"github.com/gin-gonic/gin"
)

func Start(ctx *gin.Context, dbOra *sql.DB) {
	stamp := time.Now().Format(("20060102150405"))
	log.Printf("%v - cliente/get01 - Start", stamp)
	oraRepo := repository.NewRepoOracle(dbOra)
	v := view.NewView(ctx)
	p := presenter.NewPresenter(v)
	u := usecase.NewUseCase(oraRepo, p)
	err := controller.Execute(stamp, ctx, u)
	if err != nil {
		log.Printf("%v - Error: %v", stamp, err.Error())
		p.ShowError(stamp, err.Error())
	}

}
