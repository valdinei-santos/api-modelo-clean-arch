package main

import (
	"fmt"
	"log/slog"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/api/routes"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/config"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"

	"github.com/gin-gonic/gin"
)

// Main ...
func main() {
	fmt.Println("Iniciando...")

	//AllConfig := config.NewConfig()
	//fmt.Println("Carregou configuração...")

	// setup logging
	logger.InitLogger("json", slog.LevelInfo)
	fmt.Println("Iniciou Log...")

	db := config.InitDB()
	defer db.Close()
	fmt.Println("Iniciou BD...")

	router := gin.Default()
	router.SetTrustedProxies(nil)
	routes.InitRoutes(&router.RouterGroup, db)

	slog.Info("start api-modelo-clean-arch PORT:" + config.AllConfig.APIport)
	err := router.Run(":" + config.AllConfig.APIport)
	if err != nil {
		fmt.Printf("Erro ao iniciar a API na porta %v: %v", config.AllConfig.APIport, err)
		slog.Error("Erro ao iniciar a API na porta "+config.AllConfig.APIport, slog.Any("error", err))
	}
}
