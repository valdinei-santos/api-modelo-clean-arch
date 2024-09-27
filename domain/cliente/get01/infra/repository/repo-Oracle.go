package repository

import (
	"database/sql"
	"strconv"

	"github.com/valdinei-santos/api-modelo-clean-arch/config/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/entities"
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

// QueryCarregaDados ... Exemplo com chamada de Function no BD Oracle
/* func (r *RepoOracle) QueryCarregaDados(stamp, cpf string) (bool, error) {
	var stmt *sql.Stmt
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	query := `Begin :cursor := PKG_XYZ.FUNC01(:1,:2,:3,:4); End;`
	stmt, err = tx.PrepareContext(ctx, query)
	defer stmt.Close()
	if err != nil {
		logger.Error("Erro...", err)
	}
	var result string
	_, err = stmt.ExecContext(ctx, sql.Out{Dest: &result}, param1, param2, param3, param4)
	if err != nil {
		logger.Error("Erro...", err)
	}
	if result == "S" {
		tx.Commit()
		return true, nil
	}
	return false, errors.New("ERRO-API: Erro ao carregar os dados")
} */
