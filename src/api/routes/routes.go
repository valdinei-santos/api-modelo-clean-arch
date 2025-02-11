package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	clienteGet01 "github.com/valdinei-santos/api-modelo-clean-arch/src/application/cliente/get01"
	clienteGet02 "github.com/valdinei-santos/api-modelo-clean-arch/src/application/cliente/get02"
	clientePost01 "github.com/valdinei-santos/api-modelo-clean-arch/src/application/cliente/post01"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
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

	//Lista 1 cliente específico
	api.GET("/cliente/:cpf", func(c *gin.Context) {
		logger.Info("/cliente/get01/:cpf")
		clienteGet01.Start(c, db)
	})

	//Lista todos os clientes
	api.GET("/cliente", func(c *gin.Context) {
		logger.Info("/cliente/get02")
		clienteGet02.Start(c, db)
	})

	// Cria 1 cliente
	api.POST("/cliente", func(c *gin.Context) {
		logger.Info("/cliente/post01")
		clientePost01.Start(c, db)
	})

}
