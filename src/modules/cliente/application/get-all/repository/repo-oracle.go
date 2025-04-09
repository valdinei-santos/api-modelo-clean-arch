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

// QueryLoadDataCliente ...
func (r *RepoOracle) QueryLoadDataCliente(stamp, cpf string) (*entities.Cliente, error) {
	logger.Info("Entrou... CPF:"+cpf, zap.String("id", stamp), zap.String("mtd", "cliente/get01 - Repository - QueryLoadDataCliente"))
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

// QueryLoadDataTelefone ...
func (r *RepoOracle) QueryLoadDataTelefone(stamp, cpf string) ([]entities.Telefone, error) {
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
func (r *RepoOracle) QueryLoadDataTelefone2(stamp, cpf string) ([]entities.Telefone, error) {
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

// InsertCliente ...
func (r *RepoOracle) InsertCliente(stamp string, p *dto.Cliente) error {
	logger.Info("Entrou... CPF:"+p.CPF, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Repository - InsertCliente"))

	tx, err := r.db.Begin()
	if err != nil {
		logger.Fatal("Erro Fatal", err)
	}

	query1 := `insert into cliente(cpf, nm_cliente, dt_nasc) VALUES(:1, :2, :3)`
	res1, err := tx.Exec(query1, p.CPF, p.Nome, p.DtNasc)
	if err != nil {
		tx.Rollback()
		logger.Error("Erro ao inserir cliente", err, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Repository - InsertCliente"))
		return err
	}

	query2 := `insert into telefone(cpf, numero) VALUES(:1, :2)`
	for _, tel := range p.Telefones {
		_, err = tx.Exec(query2, p.CPF, tel)
		if err != nil {
			tx.Rollback()
			logger.Error("Erro ao inserir telefone", err, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Repository - InsertCliente"))
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Erro ao fazer commit", err, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Repository - InsertCliente"))
		return err
	}

	// Exemplo para verificar o número de linhas inseridas
	rowsAffected, err := res1.RowsAffected()
	if err != nil {
		logger.Error("Erro ao obter o número de linhas afetadas", err, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Repository - InsertCliente"))
		return err
	}
	logger.Info("Número de linhas inseridas: "+fmt.Sprintf("%d", rowsAffected), zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Repository - InsertCliente"))
	return nil
}

// InsertTelefone ...
func (r *RepoOracle) InsertTelefone(stamp string, t *dto.Telefone) error {
	logger.Info("Entrou... CPF:"+t.Numero, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Repository - InsertTelefone"))

	query := `insert into pessoa(cpf, numero) VALUES(:1, :2)`
	res, err := r.db.Exec(query, t.CPF, t.Numero)
	if err != nil {
		logger.Error("Erro ao executar a inserção", err, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Repository - InsertTelefone"))
		return err
	}

	// Verificar o número de linhas afetadas
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logger.Error("Erro ao obter o número de linhas afetadas", err, zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Repository - InsertTelefone"))
		return err
	}
	logger.Info("Número de linhas inseridas: "+fmt.Sprintf("%d", rowsAffected), zap.String("id", stamp), zap.String("mtd", "cliente/post01 - Repository - InsertTelefone"))
	return nil
}
