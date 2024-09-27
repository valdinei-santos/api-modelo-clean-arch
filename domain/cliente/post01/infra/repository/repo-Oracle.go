package repository

import (
	"database/sql"
	"fmt"

	"github.com/valdinei-santos/api-modelo-clean-arch/config/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01/usecase"
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

// InsertPessoa ...
func (r *RepoOracle) InsertCliente(stamp string, p *usecase.Cliente) error {
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
func (r *RepoOracle) InsertTelefone(stamp string, t *usecase.Telefone) error {
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
