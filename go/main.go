package main

import (
	"fmt"
	"os"

	"github.com/christopherfujino/finance-platform/go/data"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: server [CSV]")
		os.Exit(1)
	}

	var rows = data.Parse(os.Args[1])
	for _, row := range rows {
		fmt.Println(row.Pretty())
	}
}
