package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else if os.Args[1] == "quest8.txt" {

		data, err := ioutil.ReadFile("quest8.txt")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(string(data))
	}
}
