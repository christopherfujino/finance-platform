package server

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/christopherfujino/chrislib-go/check"
	"github.com/christopherfujino/finance-platform/go/data"
	"github.com/christopherfujino/finance-platform/go/sqlite"
)

type TemplateWrapper struct {
	Transactions []data.Transaction
	Categories []string
}

func Serve(db *sqlite.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		var transactions = check.Two(db.GetTransactions())
		templateBytes, err := os.ReadFile("./server/index.template.html")
		if err != nil {
			panic(err)
		}
		tmpl, err := template.New("body").Parse(string(templateBytes))
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(writer, TemplateWrapper{
			Transactions: transactions,
			Categories: []string{"Uncategorized", "Eating out", "Groceries"},
		})
		if err != nil {
			panic(err)
		}
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		log.Printf("Received a request %s from %s\n", req.URL.String(), req.RemoteAddr)
	})
	log.Println("Listening on 0.0.0.0:8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic(err)
	}
}
