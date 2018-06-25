package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fiskeben/katas/bank-ocr/pkg"
)

// Run takes the path to the file and executes the command
func Run(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Unable to read file: %s\n", err.Error())
		os.Exit(1)
	}
	r := pkg.MakeReport(string(bytes))

	fmt.Println(r)
}
