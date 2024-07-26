package repository

import (
	"database/sql"
	"log"

	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01/usecase"
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
func (r *RepoOracle) InsertPessoa(stamp string, p *usecase.Pessoa) error {
	log.Printf("cliente/get01 - repoOracle - InsertPessoa cpf:%v ", p.CPF)

	query := `insert into pessoa(cpf, nm_pessoa, dt_nasc) VALUES(:1, :2, :3)`
	res, err := r.db.Exec(query, p.Nome, p.DtNasc, p.CPF)
	if err != nil {
		log.Printf("Erro ao executar a inserção: %v", err)
		return err
	}

	// Verificar o número de linhas afetadas
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Erro ao obter o número de linhas afetadas: %v", err)
		return err
	}
	log.Printf("Número de linhas inseridas: %d\n", rowsAffected)
	return nil
}

// InsertTelefone ...
func (r *RepoOracle) InsertTelefone(stamp string, t *usecase.Telefone) error {
	log.Printf("cliente/get01 - repoOracle - InsertTelefone cpf:%v ", t.Numero)

	query := `insert into pessoa(cpf, numero) VALUES(:1, :2)`
	res, err := r.db.Exec(query, t.CPF, t.Numero)
	if err != nil {
		log.Printf("Erro ao executar a inserção: %v", err)
		return err
	}

	// Verificar o número de linhas afetadas
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Erro ao obter o número de linhas afetadas: %v", err)
		return err
	}
	log.Printf("Número de linhas inseridas: %d\n", rowsAffected)
	return nil
}
