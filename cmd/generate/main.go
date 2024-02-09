package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
)

var (
	maxDigits   = flag.Int("max-digits", 3, "Max number of digits per number")
	maxDecimals = flag.Int("max-decimals", 2, "Max number of decimal places")
	minSize     = flag.Int("min-size", 2*1024, "Min file size in bytes")
)

func main() {
	flag.Parse()

	c := generateNumber()

	for c < *minSize {
		c += generateTerm()
	}
}

// No / here to avoid division by zero
var operators = []byte{'+', '-', '*'}

func generateTerm() int {
	op := rand.IntN(len(operators))

	fmt.Printf(" %c ", operators[op])

	if rand.Float32() < 0.2 {
		return 3 + generateParenExpr()
	}

	return 3 + generateNumber()
}

func generateParenExpr() int {
	fmt.Print("(")

	c := generateNumber()
	c += generateTerm()

	fmt.Print(")")

	return 2 + c
}

func generateNumber() int {
	numberOfChars := rand.IntN(*maxDigits-1) + 1

	for i := 0; i < numberOfChars; i++ {
		d := rand.IntN(9)
		if i == 0 && d == 0 {
			d = 1
		}

		fmt.Printf("%c", byte(d)+'0')
	}

	if rand.Float32() > 0.5 {
		numberOfChars++
		fmt.Print(".")

		numberOfDecimals := rand.IntN(*maxDecimals-1) + 1
		for i := 0; i < numberOfDecimals; i++ {
			d := rand.IntN(9)
			fmt.Printf("%c", byte(d)+'0')
		}

		numberOfChars += numberOfDecimals + 1
	}

	fmt.Print(' ')

	return numberOfChars + 1
}
