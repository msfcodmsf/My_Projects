package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1] //satırı, programın ilk argümanını str değişkenine atar.
	var fileType string
	if len(os.Args) < 3 { //koşulu, programın en az iki argümanla çağrılıp çağrılmadığını kontrol eder.
		fileType = "standard"
	} else {
		fileType = os.Args[2] //Program en az iki argümanla çağrıldığında, ikinci argüman (os.Args[2]) fileType değişkenine atanır. Bu, belirtilen dosyanın ismini içerir.
	}

	if str == "" {
		return
	} else if str == "\\n" {
		fmt.Println()
		return
	} else { //üstteki "\n" ile fonksiyon içindeki birleştiğinden else yapıldı.
		words := strings.Split(str, "\\n")
		for k, r := range words {
			var result []string //Boş kelimeler atlanır.
			if r == "" {        //Her harf için, ConvertText fonksiyonu kullanılarak ASCII karakteri bulunur.
				if k < len(words) { //ASCII karakterleri bir result dizisine eklenir.
					fmt.Println()
				}
				continue
			}
			for _, c := range r { // Bu döngü, r dizisindeki her karakteri (c) işler.
				result = append(result, ConvertText(fileType, (int(c)-32)*9+2, (int(c)-32)*9+9)...)
			} // Bu formüller, ConvertText fonksiyonunun argümanları olarak kullanılan sayıları üretir.

			for i := 0; i < 8; i++ { //Bu döngü, result dizisini 8x8'lik bir tabloda gösterir.
				for j := 0; j < len(result)/8; j++ {
					fmt.Print(result[i+j*8], " ")
				}
				fmt.Println()
			}
		}
	}
}

func ConvertText(file string, startLine int, endLine int) []string { //ConvertText fonksiyonu, bir metin dosyasından belirli satırları okur ve bir dizi olarak döndürür.
	//Bu fonksiyon, ASCII karakterleri oluşturmak için kullanılır.
	contentFile, err := os.Open(file + ".txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func(contentFile *os.File) { // ConvertText fonksiyonu her tamamlandığında çalışacak bir anonim fonksiyonu erteler.
		err := contentFile.Close()
		if err != nil { // Bu blok, err değişkeninin boş olmadığını (yani bir hata meydana geldiğini) kontrol eder.
			fmt.Println(err) // Eğer hata oluştuysa, hata mesajı (err) fmt.Println kullanılarak ekrana yazdırılır.
		}
	}(contentFile)
	scanner := bufio.NewScanner(contentFile) // Bu, metin dosyasını satır satır okumak için kullanılan bir tarayıcıdır.

	rowNumber := 0
	var contentArr []string // Bu dizi, okunan satırları tutacaktır.
	for scanner.Scan() {    //  döngüsü, metin dosyasında hala okunacak satırlar olduğu sürece çalışır.
		rowNumber++
		if rowNumber >= startLine && rowNumber <= endLine {
			contentArr = append(contentArr, scanner.Text()) // çağrılarak şu an okunan satır metni alınır.
		}
		if rowNumber > endLine {
			break // bitiş satırını geçerse, döngü durdurulur (break).
		}
	}

	return contentArr
}
