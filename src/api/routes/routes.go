package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	clienteCreate "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/create"
	clienteGet "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/get"
	clienteGetAll "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/get-all"
)

func InitRoutes(r *gin.RouterGroup, db *sql.DB) {

	//router.Use(cors.Default())
	r.Use(func(c *gin.Context) {
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

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := r.Group("/api/modelo")

	//Lista 1 cliente03 específico
	api.GET("/cliente03/:cpf", func(c *gin.Context) {
		logger.Info("/cliente03/get/:cpf")
		clienteGet.Start(c, db)
	})

	//Lista todos os clientes
	api.GET("/cliente03", func(c *gin.Context) {
		logger.Info("/cliente03/get")
		clienteGetAll.Start(c, db)
	})

	// Cria 1 cliente03
	api.POST("/cliente03", func(c *gin.Context) {
		logger.Info("/cliente03/post")
		clienteCreate.Start(c, db)
	})

}
