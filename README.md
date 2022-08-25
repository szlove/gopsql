# gopsql

## Installation
`$ go get github.com/szlove/gopsql`

## Usage
```
package main

import (
	"github.com/szlove/gopsql"
)

func main() {
	// DB Connection
	url := gopsql.ConnectionURL{
		host: "localhost",
		port: "5432",
		user: "example",
		password: "example1234",
		dbname:   "exampleDB",
		sslmode:  "disable",
	}
	psql, err := gopsql.Conn(url)
	if err != nil {
		panic(err)
	}
	defer psql.Close()
	
	// Transaction
	tx, err := gopsql.NewTransaction(nil)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	newBook := &Book{"my book"})
	if err := newBook.Create(); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}

type Book struct {
	name string
}

const createBook = `INSERT INTO books (name) VALUES ($1);`

func (b *Book) Create(t *gopsql.Transaction, newBook &Book) error {
	return t.Tx.ExecContext(t.Ctx, createBook, newBook.Name)
}
```
