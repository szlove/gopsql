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
	tx, err := gopsql.NewTransaction()
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

func (b *Book) Create(tx *gopsql.Transaction, newBook &Book) error {
	err := tx.ExecContext(tx.Ctx, newBook.Name)
	return err
}
```
