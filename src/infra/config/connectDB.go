package config

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/godror/godror"
)

var db *sql.DB
var dbTest *sql.DB

// InitDB - Inicia a conexao com o BD.
func InitDB() *sql.DB {
	//c := NewConfig()
	dataURL := fmt.Sprintf("user=%s password=%s connectString=%s libDir=%s",
		AllConfig.ORACLEuser, AllConfig.ORACLEpassword, AllConfig.ORACLEhost+":"+AllConfig.ORACLEport+"/"+AllConfig.ORACLEservice, AllConfig.ORACLElibdir)
	fmt.Println(dataURL)

	db, err := sql.Open("godror", dataURL)
	if err != nil {
		fmt.Println(err.Error())
		slog.Error("Erro Fatal", slog.Any("error", err), slog.String("method", "config.InitDB"))
	}
	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
		slog.Error("Erro Fatal", slog.Any("error", err), slog.String("method", "config.InitDB"))
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
func InitDBTest() *sql.DB {
	//c := NewConfig()
	dataURL := fmt.Sprintf("user=%s password=%s connectString=%s libDir=%s",
		AllConfig.TESTEuser, AllConfig.TESTEpassword, AllConfig.TESTEhost+":"+AllConfig.TESTEport+"/"+AllConfig.TESTEservice, AllConfig.TESTElibdir)
	//fmt.Println(dataURL)

	dbTest, err := sql.Open("godror", dataURL)
	if err != nil {
		fmt.Println("ERRO-DB: " + err.Error())
		slog.Error("Erro Fatal", slog.Any("error", err), slog.String("mtd", "config.InitDB"))
	}
	if err = dbTest.Ping(); err != nil {
		fmt.Println("ERRO-DB: " + err.Error())
		slog.Error("Erro Fatal", slog.Any("error", err), slog.String("mtd", "config.InitDB"))
	}
	dbTest.SetMaxOpenConns(2)
	dbTest.SetMaxIdleConns(1)
	return dbTest
}

// CloseDBTest - Fecha a conexao com o DB de TESTE.
func CloseDBTest() {
	dbTest.Close()
}
