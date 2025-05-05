package main

import (
	"fmt"
	"io"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/api/routes"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/config"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/database"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"

	"github.com/gin-gonic/gin"
)

// Main ...
func main() {
	fmt.Println("Iniciando...")

	//AllConfig := config.NewConfig()
	//fmt.Println("Carregou configuração...")

	// setup logging
	//logger.InitLogger("json", slog.LevelInfo)
	//handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	log := logger.NewSlogLogger()
	fmt.Println("Iniciou Log...")

	db := database.InitDB(log)
	defer db.Close()
	fmt.Println("Iniciou BD...")

	gin.DefaultWriter = io.Discard // Desabilita o log padrão do gin jogando para o io.Discard
	router := gin.Default()
	router.SetTrustedProxies(nil)
	routes.InitRoutes(&router.RouterGroup, db, log)

	log.Info("start api-modelo-clean-arch PORT:" + config.AllConfig.APIport)
	err := router.Run(":" + config.AllConfig.APIport)
	if err != nil {
		fmt.Printf("Erro ao iniciar a API na porta %v: %v", config.AllConfig.APIport, err)
		log.Error("Erro ao iniciar a API na porta " + config.AllConfig.APIport + " - " + err.Error())
	}
}
