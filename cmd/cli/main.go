package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	showHelp := flag.Bool("help", false, "Display help information")
	flag.Parse()

	if *showHelp {
		helpHandler()
		os.Exit(0)
	}

	if err := parse(); err != nil {
		fmt.Fprintf(os.Stderr, "Command Failed: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
