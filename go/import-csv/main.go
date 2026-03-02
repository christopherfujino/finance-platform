package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/christopherfujino/chrislib-go/check"
	_ "modernc.org/sqlite"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: import-csv [CSV]\n")
		os.Exit(1)
	}
	fmt.Printf("%v\n", os.Args)

	db := check.Two(sql.Open("sqlite", "csv.db"))

	check.Two(db.Query(`CREATE TABLE IF NOT EXISTS transactions (
   id INTEGER PRIMARY KEY,
   name TEXT NOT NULL
);`))

	// Insert
	check.Two(db.Query(`INSERT INTO transactions (name) VALUES ('Foo');`))

	var rows *sql.Rows = check.Two(db.Query("SELECT * FROM transactions;"))
	var s1, s2 string
	for rows.Next() {
		check.One(rows.Scan(&s1, &s2))
		fmt.Printf("%v %v\n", s1, s2)
	}
}
