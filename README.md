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
	t, err := gopsql.NewTransaction(nil)
	if err != nil {
		panic(err)
	}
	defer t.Rollback()
	newBook := &Book{Name: "my book"})
	if err := newBook.Create(t); err != nil {
		panic(err)
	}
	if err := t.Commit(); err != nil {
		panic(err)
	}
}

type Book struct {
	Name string
}

const createBook = `INSERT INTO books (name) VALUES ($1);`

func (b *Book) Create(t *gopsql.Transaction) error {
	_, err := t.Tx.ExecContext(t.Ctx, createBook, b.Name)
	return err
}
```
