package main

import (
	"github.com/valdinei-santos/api-modelo-clean-arch/config"
	clienteGet01 "github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01"
	clientePost01 "github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

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
	api := router.Group("/api")
	modelo := api.Group("/modelo")

	modelo.GET("/cliente/:cpf", func(c *gin.Context) {
		log.Printf("/cliente/get01/:cpf")
		clienteGet01.Start(c, db)
	})

	modelo.POST("/cliente", func(c *gin.Context) {
		log.Printf("/cliente/post01")
		clientePost01.Start(c, db)
	})

	log.Println("start github.com/valdinei-santos/api-modelo-clean-arch - PORT:", AllConfig.APIport)
	router.Run(":" + AllConfig.APIport)

}
