package main

import (
	"fmt"
	"os"

	"github.com/christopherfujino/chrislib-go/check"
	"github.com/christopherfujino/finance-platform/go/sqlite"
	"github.com/christopherfujino/finance-platform/go/server"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: server [db.sqlite]")
		os.Exit(1)
	}

	var db = check.Two(sqlite.Init(os.Args[1]))
	server.Serve(db)
}
