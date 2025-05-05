package database

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/config"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
)

var db *sql.DB
var dbTest *sql.DB

// InitDB - Inicia a conexao com o BD.
func InitDB(log logger.Logger) *sql.DB {
	//c := NewConfig()
	dataURL := fmt.Sprintf("user=%s password=%s connectString=%s libDir=%s",
		config.AllConfig.ORACLEuser,
		config.AllConfig.ORACLEpassword,
		config.AllConfig.ORACLEhost+":"+config.AllConfig.ORACLEport+"/"+config.AllConfig.ORACLEservice,
		config.AllConfig.ORACLElibdir)
	fmt.Println(dataURL)

	db, err := sql.Open("godror", dataURL)
	if err != nil {
		fmt.Println("ERRO-DB: " + err.Error())
		log.Error(err.Error(), "mtd", "sql.Open")
	}
	if err = db.Ping(); err != nil {
		fmt.Println("ERRO-DB: " + err.Error())
		log.Error(err.Error(), "mtd", "db.Ping")
	}
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(1)
	return db
}

// CloseDB - Fecha a conexao com o DB.
func CloseDB() {
	db.Close()
}

// InitDBTest - Inicia a conexao com o BD de TESTE.
func InitDBTest(log logger.Logger) *sql.DB {
	//c := NewConfig()
	dataURL := fmt.Sprintf("user=%s password=%s connectString=%s libDir=%s",
		config.AllConfig.TESTEuser,
		config.AllConfig.TESTEpassword,
		config.AllConfig.TESTEhost+":"+config.AllConfig.TESTEport+"/"+config.AllConfig.TESTEservice,
		config.AllConfig.TESTElibdir)
	//fmt.Println(dataURL)

	dbTest, err := sql.Open("godror", dataURL)
	if err != nil {
		fmt.Println("ERRO-DB-TEST: " + err.Error())
		log.Error(err.Error(), "mtd", "sql.Open")
	}
	if err = dbTest.Ping(); err != nil {
		fmt.Println("ERRO-DB-TEST: " + err.Error())
		log.Error(err.Error(), "mtd", "dbTest.Ping")
	}
	dbTest.SetMaxOpenConns(2)
	dbTest.SetMaxIdleConns(1)
	return dbTest
}

// CloseDBTest - Fecha a conexao com o DB de TESTE.
func CloseDBTest() {
	dbTest.Close()
}
