package main

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/api/router"
	"github.com/valdinei-santos/api-modelo-clean-arch/config"

	"fmt"
)

// @title           API-MODELO-CLEAN-ARCH
// @version         1.0
// @description     API com implementação usando Clean Architeture
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      site.com.br:8888
// @BasePath  /api/modelo/

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	fmt.Println("Iniciando...")

	AllConfig := config.NewConfig()
	fmt.Println("Carregou configuração...")

	// setup logging
	config.InitLogz()
	fmt.Println("Iniciou Logz...")

	db := config.InitDB()
	defer db.Close()
	fmt.Println("Iniciou BD...")

	// Initialize Router
	router.Initialize(AllConfig, db)

}
