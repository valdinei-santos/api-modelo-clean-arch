package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
)

// SqlTransaction implements ITransaction
type SqlTransaction struct {
	tx *sql.Tx
}

func (st *SqlTransaction) Commit() error {
	if st.tx == nil {
		return errors.New("transação não inicializada")
	}
	return st.tx.Commit()
}

func (st *SqlTransaction) Rollback() error {
	if st.tx == nil {
		return errors.New("transação não inicializada")
	}
	return st.tx.Rollback()
}

// RepoOracle implements Repository
type RepoOracle struct {
	db  *sql.DB
	log logger.Logger
}

// NewOraRepo create new repository
func NewRepoOracle(db *sql.DB, l logger.Logger) *RepoOracle {
	return &RepoOracle{
		db:  db,
		log: l,
	}
}

// func (r *RepoOracle) BeginTransaction() (*sql.Tx, error) {
func (r *RepoOracle) BeginTransaction() (ITransaction, error) {
	r.log.Debug("Entrou repository.BeginTransaction")
	//tx, err := r.db.Begin()
	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		r.log.Error(err.Error(), "mtd", "db.Begin")
		return nil, err
	}
	//return tx, nil
	return &SqlTransaction{tx: tx}, nil
}

// Save ...
func (r *RepoOracle) Save(p *dto.Cliente) error {
	r.log.Debug("Entrou repository.Save - CPF:" + p.CPF)
	tx, err := r.db.Begin()
	if err != nil {
		r.log.Error(err.Error(), "mtd", "db.Begin")
	}

	query1 := `insert into cliente(cpf, nm_cliente, dt_nasc) VALUES(:1, :2, :3)`
	res1, err := tx.Exec(query1, p.CPF, p.Nome, p.DtNasc)
	if err != nil {
		tx.Rollback()
		r.log.Error(err.Error(), "mtd", "tx.Exec")
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		r.log.Error(err.Error(), "mtd", "tx.Commit")
		return err
	}

	// Exemplo para verificar o número de linhas inseridas
	rowsAffected, err := res1.RowsAffected()
	if err != nil {
		r.log.Error(err.Error(), "mtd", "res1.RowsAffected")
		return err
	}
	r.log.Debug("Número de linhas inseridas: " + fmt.Sprintf("%d", rowsAffected))
	return nil
}

// FindById ...
func (r *RepoOracle) FindById(cpf string) (*entities.Cliente, error) {
	r.log.Debug("Entrou repository.FindById - CPF:" + cpf)
	var stmt *sql.Stmt
	var err error
	stmt, err = r.db.Prepare(`
		SELECT cpf, nm_cliente, dt_nasc
		  FROM cliente
	     WHERE cpf = :1
	`)
	if err != nil {
		r.log.Error(err.Error(), "mtd", "db.Prepare")
		return nil, err
	}
	defer stmt.Close()

	var c entities.Cliente
	err = stmt.QueryRow(cpf).Scan(&c.Cpf, &c.Nome, &c.DtNasc)
	if err != nil {
		r.log.Error(err.Error(), "mtd", "stmt.QueryRow")
		return &c, err
	}
	return &c, nil
}

// FindAll
func (r *RepoOracle) FindAll() (*[]entities.Cliente, error) {
	r.log.Debug("Entrou repository.FindAll")
	queryCliente := `SELECT c.cpf, c.nm_cliente, c.dt_nasc FROM cliente c`
	rows1, err := r.db.Query(queryCliente)
	if err != nil {
		r.log.Error(err.Error(), "mtd", "db.Query")
		return nil, err
	}
	defer rows1.Close()

	clientes := []entities.Cliente{}

	for rows1.Next() {
		//fmt.Println("Entrou rows1")
		var c entities.Cliente
		if err := rows1.Scan(&c.Cpf, &c.Nome, &c.DtNasc); err != nil {
			r.log.Error(err.Error(), "mtd", "rows1.Scan")
			return nil, err
		}
		clientes = append(clientes, c)
	}
	return &clientes, nil
}
