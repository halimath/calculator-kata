package main

import (
	"fmt"
	"os"

	"github.com/halimath/calc/calculator"
)

func main() {

	result, err := calculator.Eval(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: failed to evaluate: %s\n", os.Args[0], err)
		os.Exit(1)
	}

	fmt.Printf("%.5f\n", result)
}
