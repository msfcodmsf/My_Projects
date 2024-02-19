package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programParams := os.Args[1:]
	reversedParams := reverseSlice(programParams)
	for _, arg := range reversedParams {
		runes := []rune(arg)
		for _, r := range runes {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

func reverseSlice(slice []string) []string {
	// Dilimi tersine çevir
	for i := 0; i < len(slice)/2; i++ {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
	return slice
}
