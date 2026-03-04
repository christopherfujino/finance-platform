package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/christopherfujino/finance-platform/go/data"
)

type T struct {
	db *sql.DB
}

func Init(path string) (*T, error) {
	db, err := (sql.Open("sqlite", path))
	if err != nil {
		return nil, err
	}
	var d = T{db: db}

	_, err = d.db.Query(`CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY,
		date INTEGER NOT NULL,
		account TEXT NOT NULL,
		payee TEXT NOT NULL,
		amount REAL NOT NULL,
		category INTEGER NOT NULL
	);`)

	if err != nil {
		return nil, err
	}

	return &d, nil
}

func escape(s string) string {
	var b = strings.Contains(s, "'")
	if !b {
		return fmt.Sprintf("'%s'", s)
	}
	return fmt.Sprintf("'%s'", strings.ReplaceAll(s, "'", "''"))
}

func (d *T) InsertTransaction(transaction data.Transaction) error {
	// Ignore transaction.Id
	// TODO: check for ' in strings
	var query = fmt.Sprintf(
		"INSERT INTO transactions (account, date, payee, amount, category) VALUES (%s, unixepoch(%s), %s, %f, %d);",
		escape(transaction.Account),
		escape(transaction.Date.Format("2006-01-02 15:04:05")),
		escape(transaction.Payee),
		transaction.Amount,
		transaction.Category,
	)
	_, err := d.db.Query(query)
	return err
}

func (d *T) Query(query string) (*sql.Rows, error) {
	return d.db.Query(query)
}
