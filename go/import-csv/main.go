package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/christopherfujino/chrislib-go/check"
	"github.com/christopherfujino/finance-platform/go/data"
	_ "modernc.org/sqlite"
)

type DB struct {
	db *sql.DB
}

func (d DB) init(path string) *DB {
	d.db = check.Two(sql.Open("sqlite", path))

	check.Two(d.db.Query(`CREATE TABLE IF NOT EXISTS transactions (
   id INTEGER PRIMARY KEY,
   name TEXT NOT NULL
);`))

	return &d
}

func (d *DB) insertTransaction(t data.Transaction) {
	var query = fmt.Sprintf("INSERT INTO transactions (name) VALUES ('%s');", t.Account)
	d.db.Query(query)
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: import-csv [CSV]\n")
		os.Exit(1)
	}
	fmt.Printf("%v\n", os.Args)

	var db = DB{}.init("db.sqlite")

	var transactions = data.Parse(args[1])

	for _, transaction := range transactions {
		// Insert
		db.insertTransaction(transaction)
	}
	var rows *sql.Rows = check.Two(db.db.Query("SELECT * FROM transactions;"))
	var s1, s2 string
	for rows.Next() {
		check.One(rows.Scan(&s1, &s2))
		fmt.Printf("%v %v\n", s1, s2)
	}
}
