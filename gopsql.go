package gopsql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

var (
	postgres *sql.DB
)

type ConnectionURL struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (c *ConnectionURL) gen() string {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
	return url
}

func Conn(connectionURL *ConnectionURL) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionURL.gen())
	if err != nil {
		return nil, errors.Wrap(err, "sql.Open()")
	}
	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "db.Ping()")
	}
	postgres = db
	return db, nil
}

type Transaction struct {
	Tx  *sql.Tx
	Ctx context.Context
}

func NewTransaction(opts *sql.TxOptions) (*Transaction, error) {
	ctx := context.Background()
	tx, err := postgres.BeginTx(ctx, opts)
	return &Transaction{tx, ctx}, err
}

func (t *Transaction) Rollback() error { return t.Tx.Rollback() }
func (t *Transaction) Commit() error   { return t.Tx.Commit() }
