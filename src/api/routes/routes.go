package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	cliente "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases"
	telefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/application/usecases"
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

	// CLIENTES
	// Cria 1 cliente
	api.POST("/cliente", func(c *gin.Context) {
		logger.Info("POST /cliente")
		cliente.StartCreate(c, db)
	})

	// Cria 1 cliente específico com telefones
	api.POST("/cliente/telefones", func(c *gin.Context) {
		logger.Info("POST /cliente/telefones")
		cliente.StartCreateComTelefone(c, db)
	})

	//Lista 1 cliente específico
	api.GET("/cliente/:cpf", func(c *gin.Context) {
		logger.Info("GET /cliente/:cpf")
		cliente.StartGet(c, db)
	})

	//Lista todos os clientes
	api.GET("/cliente", func(c *gin.Context) {
		logger.Info("GET /cliente")
		cliente.StartGetAll(c, db)
	})

	// Lista 1 cliente específico com telefones
	api.GET("/cliente/:cpf/telefones", func(c *gin.Context) {
		logger.Info("/cliente/:cpf/telefones")
		cliente.StartGetComTelefone(c, db)
	})

	// Lista todos os clientes com telefones
	api.GET("/cliente/telefones", func(c *gin.Context) {
		logger.Info("/cliente/telefones")
		cliente.StartGetAllComTelefone(c, db)
	})

	// TELEFONES
	//Lista todos os telefones de um cliente
	api.GET("/telefone/:cpf", func(c *gin.Context) {
		logger.Info("routes.go - /telefone/:cpf")
		telefone.StartGetAll(c, db)
	})

	// Cria 1 telefone
	api.POST("/telefone", func(c *gin.Context) {
		logger.Info("routes.go - /telefone")
		telefone.StartCreate(c, db)
	})

}
