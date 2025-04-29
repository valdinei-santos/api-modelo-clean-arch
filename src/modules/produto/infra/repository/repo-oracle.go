package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
)

// RepoOracle implements Repository
type RepoOracle struct {
	db  *sql.DB
	log logger.Logger
}

// NewRepoOracle creates a new repository
func NewRepoOracle(db *sql.DB, l logger.Logger) *RepoOracle {
	return &RepoOracle{
		db:  db,
		log: l,
	}
}

func (r *RepoOracle) BeginTransaction() (*sql.Tx, error) {
	r.log.Debug("Entrou repository.BeginTransaction")
	tx, err := r.db.Begin()
	if err != nil {
		r.log.Error(err.Error(), "mtd", "db.Begin")
		return nil, err
	}
	return tx, nil
}

// Save - Salva um produto no banco de dados
func (r *RepoOracle) Save(p *dto.ProdutoDTO) error {
	r.log.Debug("Entrou repository.Save - Nome:" + p.Nome)

	tx, err := r.db.Begin()
	if err != nil {
		r.log.Error(err.Error(), "mtd", "db.Begin")
	}

	query := `INSERT INTO produto (nome, descricao, preco, qtd_estoque, categoria, fl_ativo, data_criacao, data_atualizacao)
              VALUES (:1, :2, :3, :4, :5, :6, SYSDATE, SYSDATE)`
	res, err := tx.Exec(query, p.Nome, p.Descricao, p.Preco, p.QtdEstoque, p.Categoria, p.FlAtivo)
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

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		r.log.Error(err.Error(), "mtd", "res.RowsAffected")
		return err
	}
	r.log.Debug("Número de linhas inseridas: " + fmt.Sprintf("%d", rowsAffected))
	return nil
}

// FindById - Busca um produto pelo ID
func (r *RepoOracle) FindById(id int) (*entities.Produto, error) {
	r.log.Debug("Entrou repository.FindById - ID:" + strconv.Itoa(id))
	stmt, err := r.db.Prepare(`
        SELECT id, nome, descricao, preco, qtd_estoque, categoria, fl_ativo, data_criacao, data_atualizacao
          FROM produto
         WHERE id = :1
    `)
	if err != nil {
		r.log.Error(err.Error(), "mtd", "db.Prepare")
		return nil, err
	}
	defer stmt.Close()

	var p entities.Produto
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Nome, &p.Descricao, &p.Preco, &p.QtdEstoque, &p.Categoria, &p.FlAtivo, &p.DataCriacao, &p.DataAtualizacao)
	if err != nil {
		r.log.Error(err.Error(), "mtd", "stmt.QueryRow")
		return nil, err
	}
	return &p, nil
}

// FindAll - Busca todos os produtos
func (r *RepoOracle) FindAll() (*[]entities.Produto, error) {
	r.log.Debug("Entrou repository.FindAll")
	query := `SELECT id, nome, descricao, preco, qtd_estoque, categoria, fl_ativo, data_criacao, data_atualizacao FROM produto`
	rows, err := r.db.Query(query)
	if err != nil {
		r.log.Error(err.Error(), "mtd", "db.Query")
		return nil, err
	}
	defer rows.Close()

	produtos := []entities.Produto{}
	for rows.Next() {
		var p entities.Produto
		if err := rows.Scan(&p.ID, &p.Nome, &p.Descricao, &p.Preco, &p.QtdEstoque, &p.Categoria, &p.FlAtivo, &p.DataCriacao, &p.DataAtualizacao); err != nil {
			r.log.Error(err.Error(), "mtd", "stmt.QueryRow")
			return nil, err
		}
		produtos = append(produtos, p)
	}
	return &produtos, nil
}
