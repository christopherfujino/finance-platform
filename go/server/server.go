package server

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/christopherfujino/finance-platform/go/data"
)

type TemplateWrapper struct {
	Transactions []data.Transaction
	Categories []string
}

func Serve(transactions []data.Transaction) {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		// TODO: cache this in prod
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
	log.Println("Listening on 127.0.0.1:8080")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
}
