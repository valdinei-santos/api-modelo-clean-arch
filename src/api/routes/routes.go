package routes

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	//"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/slog"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	cliente "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente"
	produto "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto"
	telefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone"
)

func InitRoutes(r *gin.RouterGroup, db *sql.DB, l logger.Logger) {

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
		log := initIdReqInLog(l, "ping")
		log.Info("### " + c.Request.Method + " " + c.Request.URL.Path)
		c.String(200, "pong")
	})

	api := r.Group("/api/modelo")

	// CLIENTES
	// Cria 1 cliente
	api.POST("/cliente", func(c *gin.Context) {
		log := initIdReqInLog(l, "cliente")
		log.Info("### " + c.Request.Method + " " + c.Request.URL.Path)
		cliente.Create(log, c, db)
	})

	// Cria 1 cliente específico com telefones
	api.POST("/cliente/telefones", func(c *gin.Context) {
		log := initIdReqInLog(l, "cliente")
		log.Info("### " + c.Request.Method + " " + c.Request.URL.Path)
		cliente.CreateComTelefone(log, c, db)
	})

	//Lista 1 cliente específico
	api.GET("/cliente/:cpf", func(c *gin.Context) {
		log := initIdReqInLog(l, "cliente")
		log.Info("### " + c.Request.Method + " " + c.Request.URL.Path)
		cliente.Get(log, c, db)
	})

	//Lista todos os clientes
	api.GET("/cliente", func(c *gin.Context) {
		log := initIdReqInLog(l, "cliente")
		log.Info("### " + c.Request.Method + " " + c.Request.URL.Path)
		cliente.GetAll(log, c, db)
	})

	// Lista 1 cliente específico com telefones
	api.GET("/cliente/:cpf/telefones", func(c *gin.Context) {
		log := initIdReqInLog(l, "cliente")
		log.Info("### " + c.Request.Method + " " + c.Request.URL.Path)
		cliente.GetComTelefone(log, c, db)
	})

	// Lista todos os clientes com telefones
	api.GET("/cliente/telefones", func(c *gin.Context) {
		log := initIdReqInLog(l, "cliente")
		log.Info("### " + c.Request.Method + " " + c.Request.URL.Path)
		cliente.GetAllComTelefone(log, c, db)
	})

	// TELEFONES
	//Lista todos os telefones de um cliente
	api.GET("/telefone/:cpf", func(c *gin.Context) {
		log := initIdReqInLog(l, "telefone")
		log.Info("### " + c.Request.Method + " " + c.Request.URL.Path)
		telefone.GetAll(log, c, db)
	})

	// Cria 1 telefone
	api.POST("/telefone", func(c *gin.Context) {
		log := initIdReqInLog(l, "telefone")
		log.Info("### " + c.Request.Method + " " + c.Request.URL.Path)
		telefone.Create(log, c, db)
	})

	// PRODUTOS
	//Lista todos os produtos
	api.GET("/produto", func(c *gin.Context) {
		log := initIdReqInLog(l, "produto")
		log.Info("### " + c.Request.Method + " " + c.Request.URL.Path)
		produto.GetAll(log, c, db)
	})

	//Lista 1 produto
	api.GET("/produto/:id", func(c *gin.Context) {
		log := initIdReqInLog(l, "produto")
		log.Info("### " + c.Request.Method + " " + c.Request.URL.Path)
		produto.Get(log, c, db)
	})

	// Cria 1 produto
	api.POST("/produto", func(c *gin.Context) {
		log := initIdReqInLog(l, "produto")
		log.Info("### " + c.Request.Method + " " + c.Request.URL.Path)
		produto.Create(log, c, db)
	})
}

func initIdReqInLog(log logger.Logger, nameResource string) logger.Logger {
	stamp := time.Now().Format(("20060102150405"))
	logNewParams := log.With("req_id", stamp, "resource", nameResource)
	return logNewParams
}
