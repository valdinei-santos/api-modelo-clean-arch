package repository

import (
	"database/sql"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
)

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

// Save ...
func (r *RepoOracle) Save(t *dto.Telefone) error {
	r.log.Debug("Entrou repository.Save - CPF: " + t.CPF)

	// Inicia a transação
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		r.log.Error(err.Error(), "mtd", "db.Begin")
		return err
	}

	// Prepara a declaração SQL para inserção
	stmt, err := tx.Prepare("INSERT INTO telefone (cpf, numero) VALUES (:1, :2)")
	if err != nil {
		tx.Rollback()
		r.log.Error(err.Error(), "mtd", "tx.Prepare")
		return err
	}
	defer stmt.Close()

	// Executa a declaração preparada com os dados
	_, err = stmt.Exec(t.CPF, t.Numero)
	if err != nil {
		tx.Rollback()
		r.log.Error(err.Error(), "mtd", "stmt.Exec")
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		r.log.Error(err.Error(), "mtd", "tx.commit")
		return err
	}

	return nil
}

func (r *RepoOracle) SaveAll(p []*dto.Telefone) error {
	r.log.Debug("Entrou repository.SaveAll")
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		r.log.Error(err.Error(), "mtd", "db.Begin")
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO telefone (cpf, numero) VALUES (:1, :2)")
	if err != nil {
		tx.Rollback()
		r.log.Error(err.Error(), "mtd", "tx.Prepare")
		return err
	}
	defer stmt.Close()

	for _, t := range p {
		_, err = stmt.Exec(t.CPF, t.Numero)
		if err != nil {
			tx.Rollback()
			r.log.Error(err.Error(), "mtd", "stmt.Exec")
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		r.log.Error(err.Error(), "mtd", "tx.Commit")
		return err
	}

	return nil
}

// FindAll ...
func (r *RepoOracle) FindAll(cpf string) ([]entities.Telefone, error) {
	r.log.Debug("Entrou repository.FindAll")
	var stmt *sql.Stmt
	var rows *sql.Rows
	var err error
	stmt, err = r.db.Prepare(`SELECT cpf, numero FROM telefone WHERE cpf = :1`)
	if err != nil {
		r.log.Error(err.Error(), "mtd", "db.Prepare")
		return nil, err
	}
	rows, err = stmt.Query(cpf)
	if err != nil {
		r.log.Error(err.Error(), "mtd", "stmt.Query")
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
			r.log.Error(err.Error(), "mtd", "rows.Scan")
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
