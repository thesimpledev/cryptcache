package main

import (
	"fmt"
	"os"
)

func main() {

	if err := parse(); err != nil {
		fmt.Fprintf(os.Stderr, "Command Failed: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
