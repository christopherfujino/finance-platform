package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

var expectedHeaders = []string{"Date", "Account", "Payee", "Category", "Exclusion", "Amount"}

func Parse(path string) []Transaction {
	var fileReader, err = os.Open(path)
	if err != nil {
		panic(err)
	}
	var reader = csv.NewReader(fileReader)
	rows, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	// [Date Account Payee Category Exclusion Amount]
	if !slices.Equal(rows[0], expectedHeaders) {
		fmt.Fprintf(os.Stderr, "Unexpected headers!\n\nExpected: %v\nGot: %v", expectedHeaders, rows[0])
	}

	var transactions = make([]Transaction, len(rows)-1)

	for i, row := range rows[1:] {
		var transaction *Transaction = &transactions[i]
		transaction.Raw = row
		transaction.Date, err = time.Parse("Jan 2, 2006", row[0])
		if err != nil {
			panic(err)
		}
		transaction.Account = row[1]
		transaction.Payee = row[2]
		transaction.Amount, err = strconv.ParseFloat(strings.TrimSpace(row[5]), 64)
		if err != nil {
			panic(err)
		}
	}

	return transactions
}
