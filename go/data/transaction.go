package data

import (
	"fmt"
	"time"

	"github.com/christopherfujino/finance-platform/go/data/category"
)

// Should this be a more dynamic format, to faciliate scripting?
type Transaction struct {
	Date     time.Time
	Account  string
	Payee    string
	Amount   float64
	Category category.T

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
