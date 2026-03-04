package main

import (
	"fmt"
	"os"

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

	transactions = check.Two(db.GetTransactions())
	for _, transaction := range transactions {
		fmt.Printf("%s\n", transaction.Pretty())
	}
}
