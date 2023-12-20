package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programParams := os.Args[1:]

	for _, arg := range programParams {
		runes := []rune(arg)
		for _, r := range runes {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
