package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/valdinei-santos/api-modelo-clean-arch/config"
	clienteGet01 "github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01"
	clientePost01 "github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01"

	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/valdinei-santos/api-modelo-clean-arch/api/docs"
)

func Initialize(AllConfig *config.Config, db *sql.DB) {
	// Initialize Router
	//router := gin.Default()

	// programmatically set swagger info
	// docs.SwaggerInfo.Title = "API-MODELO-CLEAN-ARCH"
	// docs.SwaggerInfo.Description = "Template de API GO com estrutura Clean Architeture."
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "site.com.br"
	// docs.SwaggerInfo.BasePath = "/v2"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	//router.Use(cors.Default())

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	basePath := "/api/modelo"
	docs.SwaggerInfo.BasePath = basePath
	api := router.Group(basePath)

	api.GET("/cliente/:cpf", GetCliente(db))
	/* api.GET("/cliente/:cpf", func(c *gin.Context) {
		log.Printf("/cliente/get01/:cpf")
		clienteGet01.Start(c, db)
	}) */

	api.POST("/cliente", func(c *gin.Context) {
		log.Printf("/cliente/post01")
		clientePost01.Start(c, db)
	})

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("start github.com/valdinei-santos/api-modelo-clean-arch - PORT:", AllConfig.APIport)
	router.Run(":" + AllConfig.APIport)
}

// @BasePath /api/modelo

// GetCliente godoc
// @Summary Get cliente by CPF
// @Description get cliente by CPF
// @ID get-cliente-by-cpf
// @Accept  json
// @Produce  json
// @Param   cpf      path   string     true  "Cliente CPF"
// @Success 200 {object} ClienteResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /cliente/get01/{cpf} [get]
func GetCliente(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("/cliente/get01/:cpf")
		clienteGet01.Start(c, db)
	}
}
