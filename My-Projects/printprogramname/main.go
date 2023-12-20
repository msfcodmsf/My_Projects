package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[0]
	nameRunes := []rune(programName)
	for _, r := range nameRunes[2:] {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
