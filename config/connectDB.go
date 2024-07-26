package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

var db *sql.DB
var dbTest *sql.DB

// InitDB - Inicia a conexao com o BD.
func InitDB() *sql.DB {
	c := NewConfig()
	dataURL := fmt.Sprintf("user=%s password=%s connectString=%s libDir=%s",
		c.ORACLEuser, c.ORACLEpassword, c.ORACLEhost+":"+c.ORACLEport+"/"+c.ORACLEservice, c.ORACLElibdir)
	fmt.Println(dataURL)

	db, err := sql.Open("godror", dataURL)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err.Error())
	}
	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
		log.Fatal(err.Error())
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
	c := NewConfig()
	dataURL := fmt.Sprintf("user=%s password=%s connectString=%s libDir=%s",
		c.TESTEuser, c.TESTEpassword, c.TESTEhost+":"+c.TESTEport+"/"+c.TESTEservice, c.TESTElibdir)
	//fmt.Println(dataURL)

	dbTest, err := sql.Open("godror", dataURL)
	if err != nil {
		fmt.Println("ERRO-DB: " + err.Error())
		log.Fatal(err.Error())
	}
	if err = dbTest.Ping(); err != nil {
		fmt.Println("ERRO-DB: " + err.Error())
		log.Fatal(err.Error())
	}
	dbTest.SetMaxOpenConns(2)
	dbTest.SetMaxIdleConns(1)
	return dbTest
}

// CloseDBTest - Fecha a conexao com o DB de TESTE.
func CloseDBTest() {
	dbTest.Close()
}
