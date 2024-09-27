package main

import (
	"fmt"

	"github.com/valdinei-santos/api-modelo-clean-arch/api/routes"
	"github.com/valdinei-santos/api-modelo-clean-arch/config"
	"github.com/valdinei-santos/api-modelo-clean-arch/config/logger"

	"github.com/gin-gonic/gin"
)

// Main ...
func main() {
	fmt.Println("Iniciando...")

	//AllConfig := config.NewConfig()
	//fmt.Println("Carregou configuração...")

	// setup logging
	//config.InitLog()
	//fmt.Println("Iniciou Log...")

	db := config.InitDB()
	defer db.Close()
	fmt.Println("Iniciou BD...")

	router := gin.Default()
	router.SetTrustedProxies(nil)
	routes.InitRoutes(&router.RouterGroup, db)

	logger.Info("start api-modelo-clean-arch PORT:" + config.AllConfig.APIport)
	err := router.Run(":" + config.AllConfig.APIport)
	if err != nil {
		fmt.Printf("Erro ao iniciar a API na porta %v: %v", config.AllConfig.APIport, err)
		logger.Error("Erro ao iniciar a API na porta "+config.AllConfig.APIport, err)
	}
}
