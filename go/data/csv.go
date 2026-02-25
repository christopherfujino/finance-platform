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

// Should this be a more dynamic format, to faciliate scripting?
type Transaction struct {
	Date    time.Time
	Account string
	Payee   string
	Amount  float64

	Raw []string
}

func (t Transaction) Pretty() string {
	var currencyString string
	if t.Amount < 0 {
		currencyString = fmt.Sprintf("-$%.2f", -t.Amount)
	} else {
		currencyString = fmt.Sprintf("$%.2f", t.Amount)
	}
	return fmt.Sprintf("%s %s: %s %s", t.Date.Format("2006-01-02"), t.Account, t.Payee, currencyString)
}

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
