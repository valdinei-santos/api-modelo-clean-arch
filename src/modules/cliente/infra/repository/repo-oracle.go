package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
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

func (r *RepoOracle) BeginTransaction(stamp string) (*sql.Tx, error) {
	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Erro ao iniciar transação", err, zap.String("id", stamp), zap.String("mtd", "cliente - Repository - BeginTransaction"))
		return nil, err
	}
	return tx, nil
}

// Save ...
func (r *RepoOracle) Save(stamp string, p *dto.Cliente) error {
	logger.Info("Entrou... CPF:"+p.CPF, zap.String("id", stamp), zap.String("mtd", "cliente - Repository - Save"))

	tx, err := r.db.Begin()
	if err != nil {
		logger.Fatal("Erro Fatal", err)
	}

	query1 := `insert into cliente(cpf, nm_cliente, dt_nasc) VALUES(:1, :2, :3)`
	res1, err := tx.Exec(query1, p.CPF, p.Nome, p.DtNasc)
	if err != nil {
		tx.Rollback()
		logger.Error("Erro ao inserir cliente", err, zap.String("id", stamp), zap.String("mtd", "cliente - Repository - Save"))
		return err
	}

	/* query2 := `insert into telefone(cpf, numero) VALUES(:1, :2)`
	for _, tel := range p.Telefones {
		_, err = tx.Exec(query2, p.CPF, tel)
		if err != nil {
			tx.Rollback()
			logger.Error("Erro ao inserir telefone", err, zap.String("id", stamp), zap.String("mtd", "cliente - Repository - Save"))
			return err
		}
	} */
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Erro ao fazer commit", err, zap.String("id", stamp), zap.String("mtd", "cliente - Repository - Save"))
		return err
	}

	// Exemplo para verificar o número de linhas inseridas
	rowsAffected, err := res1.RowsAffected()
	if err != nil {
		logger.Error("Erro ao obter o número de linhas afetadas", err, zap.String("id", stamp), zap.String("mtd", "cliente - Repository - Save"))
		return err
	}
	logger.Info("Número de linhas inseridas: "+fmt.Sprintf("%d", rowsAffected), zap.String("id", stamp), zap.String("mtd", "cliente - Repository - Save"))
	return nil
}

// FindById ...
func (r *RepoOracle) FindById(stamp, cpf string) (*entities.Cliente, error) {
	logger.Info("Entrou... CPF:"+cpf, zap.String("id", stamp), zap.String("mtd", "cliente - Repository - FindById"))
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

// FindByIdTelefone ...
func (r *RepoOracle) FindByIdTelefone(stamp, cpf string) ([]entities.Telefone, error) {
	logger.Info("Entrou... CPF:"+cpf, zap.String("id", stamp), zap.String("mtd", "cliente/get01 - Repository - QueryLoadDataTelefone"))
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
		logger.Info(t.Numero, zap.String("id", stamp), zap.String("mtd", "cliente/get01 - Repository - QueryLoadDataTelefone"))
		tels = append(tels, t)
		count++
	}
	logger.Info("Counts Telefone:"+strconv.Itoa(count), zap.String("id", stamp), zap.String("mtd", "cliente/get01 - Repository - QueryLoadDataTelefone"))
	if count == 0 {
		logger.Info("Telefone: No rows!", zap.String("id", stamp), zap.String("mtd", "cliente/get01 - Repository - QueryLoadDataTelefone"))
		t = entities.Telefone{Cpf: "", Numero: ""}
		tels = append(tels, t)
	}
	return tels, nil
}

// FindAll
func (r *RepoOracle) FindAll(stamp string) (*[]entities.Cliente, error) {
	logger.Info("Entrou...", zap.String("id", stamp), zap.String("mtd", "cliente/get02 - Repository - QueryLoadAllClientes"))
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

// FindAllTelefone ...
func (r *RepoOracle) FindAllTelefone(stamp, cpf string) ([]entities.Telefone, error) {
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
