package repository

import (
	"database/sql"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
)

// RepoOracle implements Repository
type RepoOracle struct {
	db *sql.DB
}

// NewRepoOracle creates a new repository
func NewRepoOracle(db *sql.DB) *RepoOracle {
	return &RepoOracle{
		db: db,
	}
}

func (r *RepoOracle) BeginTransaction(stamp string) (*sql.Tx, error) {
	tx, err := r.db.Begin()
	if err != nil {
		slog.Error("Erro ao iniciar transação", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "produto - Repository - BeginTransaction"))
		return nil, err
	}
	return tx, nil
}

// Save - Salva um produto no banco de dados
func (r *RepoOracle) Save(stamp string, p *dto.ProdutoDTO) error {
	slog.Info("Entrou... Nome:"+p.Nome, slog.String("id", stamp), slog.String("mtd", "produto - Repository - Save"))

	tx, err := r.db.Begin()
	if err != nil {
		slog.Error("Erro ao criar transação", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "produto - Repository - Save"))
	}

	query := `INSERT INTO produto (nome, descricao, preco, qtd_estoque, categoria, fl_ativo, data_criacao, data_atualizacao)
              VALUES (:1, :2, :3, :4, :5, :6, SYSDATE, SYSDATE)`
	res, err := tx.Exec(query, p.Nome, p.Descricao, p.Preco, p.QtdEstoque, p.Categoria, p.FlAtivo)
	if err != nil {
		tx.Rollback()
		slog.Error("Erro ao inserir produto", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "produto - Repository - Save"))
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		slog.Error("Erro ao fazer commit", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "produto - Repository - Save"))
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		slog.Error("Erro ao obter o número de linhas afetadas", slog.Any("error", err), slog.String("id", stamp), slog.String("mtd", "produto - Repository - Save"))
		return err
	}
	slog.Info("Número de linhas inseridas: "+fmt.Sprintf("%d", rowsAffected), slog.String("id", stamp), slog.String("mtd", "produto - Repository - Save"))
	return nil
}

// FindById - Busca um produto pelo ID
func (r *RepoOracle) FindById(stamp string, id int) (*entities.Produto, error) {
	slog.Info("Entrou... ID:"+strconv.Itoa(id), slog.String("id", stamp), slog.String("mtd", "produto - Repository - FindById"))
	stmt, err := r.db.Prepare(`
        SELECT id, nome, descricao, preco, qtd_estoque, categoria, fl_ativo, data_criacao, data_atualizacao
          FROM produto
         WHERE id = :1
    `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p entities.Produto
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Nome, &p.Descricao, &p.Preco, &p.QtdEstoque, &p.Categoria, &p.FlAtivo, &p.DataCriacao, &p.DataAtualizacao)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// FindAll - Busca todos os produtos
func (r *RepoOracle) FindAll(stamp string) (*[]entities.Produto, error) {
	slog.Info("Entrou...", slog.String("id", stamp), slog.String("mtd", "produto - Repository - FindAll"))
	query := `SELECT id, nome, descricao, preco, qtd_estoque, categoria, fl_ativo, data_criacao, data_atualizacao FROM produto`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	produtos := []entities.Produto{}
	for rows.Next() {
		var p entities.Produto
		if err := rows.Scan(&p.ID, &p.Nome, &p.Descricao, &p.Preco, &p.QtdEstoque, &p.Categoria, &p.FlAtivo, &p.DataCriacao, &p.DataAtualizacao); err != nil {
			return nil, err
		}
		produtos = append(produtos, p)
	}
	return &produtos, nil
}
