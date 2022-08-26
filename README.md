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
		Host: "localhost",
		Port: "5432",
		User: "example",
		Password: "example1234",
		DBName:   "exampleDB",
		SSLMode:  "disable",
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
	newBook := &Book{Name: "my book"})
	if err := newBook.Create(); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}

type Book struct {
	Name string
}

const createBook = `INSERT INTO books (name) VALUES ($1);`

func (b *Book) Create(t *gopsql.Transaction, newBook *Book) error {
	_, err := t.Tx.ExecContext(t.Ctx, createBook, newBook.Name)
	return err
}
```
