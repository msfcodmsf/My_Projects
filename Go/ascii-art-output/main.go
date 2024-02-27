package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	outputFileArg := os.Args[1][:9]
	if outputFileArg != "--output=" {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}
	outputFileName := os.Args[1][9:]
	outputFile, _ := os.Create(outputFileName)
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}
	fileType := "standard"
	str := os.Args[2]
	if len(os.Args) == 4 {
		fileType = os.Args[3]
	}
	if str == "" {
		return
	} else if str == "\\n" {
		fmt.Println()
		return
	} else {
		words := strings.Split(str, "\\n")
		for k, r := range words {
			var result []string
			if r == "" {
				if k < len(words) {
					fmt.Println()
				}
				continue
			}
			for _, c := range r {
				result = append(result, ConvertText(fileType, (int(c)-32)*9+2, (int(c)-32)*9+9)...)
			}

			for i := 0; i < 8; i++ {
				for j := 0; j < len(result)/8; j++ {
					_, err := outputFile.WriteString(result[i+j*8])

					if err != nil {
						log.Fatal(err)
						return
					}

				}

				_, err := outputFile.WriteString("\n")
				if err != nil {
					return
				}
			}
			outputFile.WriteString("\n")
		}
	}
}

func ConvertText(file string, startLine int, endLine int) []string {

	contentFile, err := os.Open(file + ".txt")
	if err != nil {
		fmt.Println("File not exist, Try with these ones:(thinkertoy, shadow, standard)")
	}

	defer func(contentFile *os.File) {
		err := contentFile.Close()
		if err != nil { // Bu blok, err değişkeninin boş olmadığını (yani bir hata meydana geldiğini) kontrol eder.
			fmt.Println(err) // Eğer hata oluştuysa, hata mesajı (err) fmt.Println kullanılarak ekrana yazdırılır.
		}
	}(contentFile)
	scanner := bufio.NewScanner(contentFile)

	rowNumber := 0
	var contentArr []string
	for scanner.Scan() {
		rowNumber++
		if rowNumber >= startLine && rowNumber <= endLine {
			contentArr = append(contentArr, scanner.Text())
		}
		if rowNumber > endLine {
			break
		}
	}

	return contentArr
}
