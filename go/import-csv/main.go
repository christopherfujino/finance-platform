package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/christopherfujino/chrislib-go/check"
	"github.com/christopherfujino/finance-platform/go/sqlite"
	"github.com/christopherfujino/finance-platform/go/data"
	_ "modernc.org/sqlite"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: import-csv [CSV]\n")
		os.Exit(1)
	}
	fmt.Printf("%v\n", os.Args)

	db := check.Two(sqlite.Init("db.sqlite"))

	var transactions = data.Parse(args[1])

	for _, transaction := range transactions {
		// Insert
		check.One(db.InsertTransaction(transaction))
	}
	var rows *sql.Rows = check.Two(db.Query("SELECT * FROM transactions;"))
	var dateInt int64
	for rows.Next() {
		var t data.Transaction
		check.One(rows.Scan(&t.Id, &dateInt, &t.Account, &t.Payee, &t.Amount, &t.Category))
		t.Date = time.Unix(dateInt, 0)
		fmt.Printf("%s\n", t.Pretty())
	}
}
