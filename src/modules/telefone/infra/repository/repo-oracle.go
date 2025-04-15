package repository

import (
	"database/sql"
	"log/slog"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
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

// Save ...
func (r *RepoOracle) Save(stamp string, t *dto.Telefone) error {
	slog.Info("Entrou... CPF: "+t.CPF, slog.String("id", stamp), slog.String("mtd", "Repository - telefone - Save"))

	// Inicia a transação
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		slog.Error("Erro ao criar a transação", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "Repository - telefone - Save"))
		return err
	}

	// Prepara a declaração SQL para inserção
	stmt, err := tx.Prepare("INSERT INTO telefone (cpf, numero) VALUES (:1, :2)")
	if err != nil {
		tx.Rollback()
		slog.Error("Erro ao preparar a inserção", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "Repository - telefone - Save"))
		return err
	}
	defer stmt.Close()

	// Executa a declaração preparada com os dados
	_, err = stmt.Exec(t.CPF, t.Numero)
	if err != nil {
		tx.Rollback()
		slog.Error("Erro ao executar a inserção", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "Repository - telefone - Save"))
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		slog.Error("Erro ao commitar a inserção", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "Repository - telefone - Save"))
		return err
	}

	return nil
}

func (r *RepoOracle) SaveAll(stamp string, p []*dto.Telefone) error {
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "Repository - telefone - SaveAll"))
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		slog.Error("Erro ao criar a transação", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "Repository - telefone - SaveAll"))
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO telefone (cpf, numero) VALUES (:1, :2)")
	if err != nil {
		tx.Rollback()
		slog.Error("Erro ao preparar a inserção", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "Repository - telefone - SaveAll"))
		return err
	}
	defer stmt.Close()

	for _, t := range p {
		_, err = stmt.Exec(t.CPF, t.Numero)
		if err != nil {
			tx.Rollback()
			slog.Error("Erro ao executar a inserção", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "Repository - telefone - SaveAll"))
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		slog.Error("Erro ao commitar a inserção", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "Repository - telefone - SaveAll"))
		return err
	}

	return nil
}

// FindAll ...
func (r *RepoOracle) FindAll(stamp, cpf string) ([]entities.Telefone, error) {
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "Repository - telefone - FindAll"))
	var stmt *sql.Stmt
	var rows *sql.Rows
	var err error
	stmt, err = r.db.Prepare(`SELECT cpf, numero FROM telefone WHERE cpf = :1`)
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
