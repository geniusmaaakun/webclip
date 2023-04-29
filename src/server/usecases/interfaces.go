package usecases

import (
	"context"
	"database/sql"
	"webclip/src/server/models"
)

//unused

//repoのinterfaceを書く

type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type TransactionRepo interface {
	NewTransaction(on bool) (Transaction, error)
	Transaction(CTx Transaction, txFunc func(*sql.Tx) (interface{}, error)) (data interface{}, err error)
}

type MarkdownRepo interface {
	Create(tx Transaction, md *models.MarkdownMemo) error
	Delete(tx Transaction, md *models.MarkdownMemo) error
	DeleteByTitle(tx Transaction, title string) error
	DeleteByPath(tx Transaction, path string) error
	Update(tx Transaction, md *models.MarkdownMemo) error
	FindById(tx Transaction, id int) (*models.MarkdownMemo, error)
	FindAll(tx Transaction) ([]*models.MarkdownMemo, error)
	FindByTitle(tx Transaction, title string) ([]*models.MarkdownMemo, error)
	FindByPath(tx Transaction, path string) ([]*models.MarkdownMemo, error)
	FindByTitleLastOne(tx Transaction, title string) (*models.MarkdownMemo, error)
	FindBySrcUrl(tx Transaction, srcUrl string) (*models.MarkdownMemo, error)
	SearchByTitle(tx Transaction, title string) ([]*models.MarkdownMemo, error)
}
