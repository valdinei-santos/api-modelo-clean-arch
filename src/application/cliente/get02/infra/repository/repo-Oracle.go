package repository

import (
	"database/sql"
	"fmt"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"go.uber.org/zap"
	//. "api-trust/infra/pkg/log"
)

// RepoOracle implements Repository
type RepoOracle struct {
	db *sql.DB
}

// NewOraRepo create new repository
func NewRepoOracle(db *sql.DB) *RepoOracle {
	return &RepoOracle{
		db: db,
	}
}

// QueryLoadAllClientes ...
func (r *RepoOracle) QueryLoadAllClientes(stamp string) (*[]entities.ClienteComTel, error) {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente/get02 - Repository - QueryLoadAllClientes"))
	queryCliente := `SELECT c.cpf, c.nm_cliente, c.dt_nasc FROM cliente c`
	rows1, err := r.db.Query(queryCliente)
	if err != nil {
		return nil, err
	}
	defer rows1.Close()

	queryTel := `SELECT t.numero FROM telefone t WHERE t.cpf = :1`
	/*var rows2 *sql.Rows
	defer rows2.Close()
	*/
	//clientes := make([]entities.ClienteComTel, 3)
	c := entities.ClienteComTel{
		Cpf:       "5",
		Nome:      "5",
		DtNasc:    "5",
		Telefones: []string{"23", "4344"},
	}
	//cs := []entities.ClienteComTel{c}
	clientes := []entities.ClienteComTel{c}
	//tels := []entities.ClienteComTel{c}

	for rows1.Next() {
		//fmt.Println("Entrou rows1")
		var c entities.ClienteComTel
		if err := rows1.Scan(&c.Cpf, &c.Nome, &c.DtNasc); err != nil {
			return nil, err
		}
		rows2, err := r.db.Query(queryTel, c.Cpf)
		if err != nil {
			return nil, err
		}
		var t string
		var tels []string
		for rows2.Next() {
			//fmt.Println("Entrou rows2")
			if err := rows2.Scan(&t); err != nil {
				return nil, err
			}
			tels = append(tels, t)
		}
		c.Telefones = tels
		//fmt.Println(clientes)
		clientes = append(clientes, c)
	}
	fmt.Println(clientes)
	return &clientes, nil
	//return cs, nil
}

// QueryLoadDataTelefone ...
func (r *RepoOracle) QueryLoadDataTelefone(stamp, cpf string) ([]entities.Telefone, error) {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente/get02 - Repository - QueryLoadDataTelefone"))
	var stmt *sql.Stmt
	var rows *sql.Rows
	var err error
	stmt, err = r.db.Prepare(`
		SELECT cpf, numero 
		  FROM telefone
	     WHERE cpf = :1
	`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(cpf)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	defer stmt.Close()
	count := 0
	var t entities.Telefone
	var tels []entities.Telefone
	for rows.Next() {
		err = rows.Scan(&t.Cpf, &t.Numero)
		if err != nil {
			return nil, err
		}
		tels = append(tels, t)
		count++
	}
	if count == 0 {
		t = entities.Telefone{Cpf: "", Numero: ""}
		tels = append(tels, t)
	}
	return tels, nil
}
