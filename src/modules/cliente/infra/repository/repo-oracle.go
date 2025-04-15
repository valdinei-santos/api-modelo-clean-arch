package repository

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
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

func (r *RepoOracle) BeginTransaction(stamp string) (*sql.Tx, error) {
	tx, err := r.db.Begin()
	if err != nil {
		slog.Error("Erro ao iniciar transação", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "cliente - Repository - BeginTransaction"))
		return nil, err
	}
	return tx, nil
}

// Save ...
func (r *RepoOracle) Save(stamp string, p *dto.Cliente) error {
	slog.Info("Entrou... CPF:"+p.CPF, slog.String("id", stamp), slog.String("mtd", "cliente - Repository - Save"))

	tx, err := r.db.Begin()
	if err != nil {
		slog.Error("Erro Fatal", slog.Any("error", err))
	}

	query1 := `insert into cliente(cpf, nm_cliente, dt_nasc) VALUES(:1, :2, :3)`
	res1, err := tx.Exec(query1, p.CPF, p.Nome, p.DtNasc)
	if err != nil {
		tx.Rollback()
		slog.Error("Erro ao inserir cliente", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "cliente - Repository - Save"))
		return err
	}

	/* query2 := `insert into telefone(cpf, numero) VALUES(:1, :2)`
	for _, tel := range p.Telefones {
		_, err = tx.Exec(query2, p.CPF, tel)
		if err != nil {
			tx.Rollback()
			slog.Error("Erro ao inserir telefone", err, slog.String("id", stamp), slog.String("mtd", "cliente - Repository - Save"))
			return err
		}
	} */
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		slog.Error("Erro ao fazer commit", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "cliente - Repository - Save"))
		return err
	}

	// Exemplo para verificar o número de linhas inseridas
	rowsAffected, err := res1.RowsAffected()
	if err != nil {
		slog.Error("Erro ao obter o número de linhas afetadas", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "cliente - Repository - Save"))
		return err
	}
	slog.Info("Número de linhas inseridas: "+fmt.Sprintf("%d", rowsAffected), slog.String("id", stamp), slog.String("mtd", "cliente - Repository - Save"))
	return nil
}

// FindById ...
func (r *RepoOracle) FindById(stamp, cpf string) (*entities.Cliente, error) {
	slog.Info("Entrou... CPF:"+cpf, slog.String("id", stamp), slog.String("mtd", "cliente - Repository - FindById"))
	var stmt *sql.Stmt
	var err error
	stmt, err = r.db.Prepare(`
		SELECT cpf, nm_cliente, dt_nasc
		  FROM cliente
	     WHERE cpf = :1
	`)
	if err != nil {
		/* if err == sql.ErrNoRows {
			// Handle the case of no rows returned.
		} */
		return nil, err
	}
	defer stmt.Close()

	var c entities.Cliente
	err = stmt.QueryRow(cpf).Scan(&c.Cpf, &c.Nome, &c.DtNasc)
	if err != nil {
		/* if err == sql.ErrNoRows {
			// Handle the case of no rows returned.
		} */
		return &c, err
	}
	return &c, nil
}

// FindAll
func (r *RepoOracle) FindAll(stamp string) (*[]entities.Cliente, error) {
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "cliente/get02 - Repository - QueryLoadAllClientes"))
	queryCliente := `SELECT c.cpf, c.nm_cliente, c.dt_nasc FROM cliente c`
	rows1, err := r.db.Query(queryCliente)
	if err != nil {
		return nil, err
	}
	defer rows1.Close()

	clientes := []entities.Cliente{}

	for rows1.Next() {
		//fmt.Println("Entrou rows1")
		var c entities.Cliente
		if err := rows1.Scan(&c.Cpf, &c.Nome, &c.DtNasc); err != nil {
			return nil, err
		}
		clientes = append(clientes, c)
	}
	fmt.Println(clientes)
	return &clientes, nil
	//return cs, nil
}
