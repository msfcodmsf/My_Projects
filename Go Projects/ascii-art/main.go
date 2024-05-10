package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Bu kod, kullanıcıdan alınan bir metni ASCII sanatına dönüştürür.
Kod, metni kelimelerine ayırır ve her kelimeyi ASCII karakterlerinden oluşan bir diziye dönüştürür.
Daha sonra bu diziyi 8x8'lik bir tabloda gösterir.
*/
func main() {
	str := os.Args[1]
	if str == "" {
		return
	} else if str == "\\n" {
		fmt.Println()
	} else { //üstteki "\n" ile fonksiyon içindeki birleştiğinden else yapıldı.
		words := strings.Split(str, "\\n") // str değişkeni, kullanıcıdan alınan metni içerir. Split fonksiyonu, metni yeni satır karakterlerine göre böler ve her kelimeyi words dizisine ekler.
		//Bu döngü, words dizisindeki her kelimeyi işler. Her kelime için:
		for k, r := range words {
			var result []string //Boş kelimeler atlanır.
			if r == "" {        //Her harf için, ConvertText fonksiyonu kullanılarak ASCII karakteri bulunur.
				if k < len(words) { //ASCII karakterleri bir result dizisine eklenir.
					fmt.Println()
				}
				continue
			}
			for _, c := range r { // Bu döngü, r dizisindeki her karakteri (c) işler.
				result = append(result, ConvertText((int(c)-32)*9+2, (int(c)-32)*9+9)...) // Bu formüller, ConvertText fonksiyonunun argümanları olarak kullanılan sayıları üretir.
			}

			for i := 0; i < 8; i++ { //Bu döngü, result dizisini 8x8'lik bir tabloda gösterir.
				for j := 0; j < len(result)/8; j++ {
					fmt.Print(result[i+j*8], " ")
				}
				fmt.Println()
			}
		}
	}
}

func ConvertText(startLine int, endLine int) []string { //ConvertText fonksiyonu, bir metin dosyasından belirli satırları okur ve bir dizi olarak döndürür. Bu fonksiyon, ASCII karakterleri oluşturmak için kullanılır.
	contentFile, err := os.Open("standard.txt") // Bu satır, os.Open fonksiyonunu kullanarak "standard.txt" dosyasını açmayı dener.
	if err != nil {                             // Eğer dosya açılırsa, dosya tanımlayıcı bir contentFile değişkenine atanır.
		fmt.Println(err) // Eğer hata oluşursa, hata bilgisi err değişkenine atanır.
	}

	defer func(contentFile *os.File) { // ConvertText fonksiyonu her tamamlandığında çalışacak bir anonim fonksiyonu erteler.
		err := contentFile.Close()
		if err != nil { // Bu blok, err değişkeninin boş olmadığını (yani bir hata meydana geldiğini) kontrol eder.
			fmt.Println(err) // Eğer hata oluştuysa, hata mesajı (err) fmt.Println kullanılarak ekrana yazdırılır.
		}
	}(contentFile)
	contentFile.Seek(0, 0)                   // Go dilinde bir dosya üzerindeki imleç konumunu değiştirmek için kullanılan bir metoddur.
	scanner := bufio.NewScanner(contentFile) // Bu, metin dosyasını satır satır okumak için kullanılan bir tarayıcıdır.

	rowNumber := 0 // Bu sayı, şu anda okunan satırın numarasını takip eder.
	/*
		Bu kod parçası, startLine'dan başlayarak endLine'a kadar olan satırları metin dosyasından okur ve bu satırları içeren bir diziyi geri döndürür.
	*/
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
