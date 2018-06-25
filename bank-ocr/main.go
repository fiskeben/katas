package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fiskeben/katas/bank-ocr/cmd"
)

func main() {
	flag.Parse()
	path := flag.Arg(0)

	if path == "" {
		fmt.Println("Usage: ocr <filename>")
		os.Exit(1)
	}

	cmd.Run(path)
}
