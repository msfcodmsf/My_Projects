package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	bubbleSort(arguments)

	for _, arg := range arguments {
		runes := []rune(arg)
		for _, r := range runes {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

func bubbleSort(arr []string) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
