package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	for _, e := range a {
		var s string
		s += string(e)

		for _, s := range e {
			z01.PrintRune(s)
		}
		z01.PrintRune('\n')
	}
}
