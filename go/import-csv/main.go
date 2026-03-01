package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: import-csv [CSV]\n")
		os.Exit(1)
	}
	fmt.Printf("%v\n", os.Args)
}
