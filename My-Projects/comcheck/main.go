package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	for _, arr := range arg {
		if arr == "01" || arr == "galaxy" || arr == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
