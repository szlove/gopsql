package gopsql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
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

var Conns = make(map[string]*sql.DB)

func Conn(connectionName string, connectionURL *ConnectionURL) error {
	db, err := sql.Open("postgres", connectionURL.gen())
	if err != nil {
		return errors.Wrap(err, "sql.Open()")
	}
	if err := db.Ping(); err != nil {
		return errors.Wrap(err, "db.Ping()")
	}
	Conns[connectionName] = db
	return nil
}

type Transaction struct {
	Tx  *sql.Tx
	Ctx context.Context
}

func NewTransaction(connectionName string, opts *sql.TxOptions) (*Transaction, error) {
	conn, ok := Conns[connectionName]
	if !ok {
		return nil, errors.New("Connection not found.")
	}
	ctx := context.Background()
	tx, err := conn.BeginTx(ctx, opts)
	return &Transaction{tx, ctx}, err
}

func (t *Transaction) Rollback() error { return t.Tx.Rollback() }
func (t *Transaction) Commit() error   { return t.Tx.Commit() }
