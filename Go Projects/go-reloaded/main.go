package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) == 3 { // os.Args değişkeninin uzunluğu 3 ile karşılaştırılır. Bu, programın en az 3 komut satırı argümanına ihtiyaç duyduğunu gösterir.
		input, err := os.ReadFile(os.Args[1]) // os.Args[1] ifadesi kullanılarak, ikinci komut satırı argümanı olarak verilen dosyanın adı alınır.
		if err != nil {
			log.Fatalf("Error reading input file: %v", err) // Eğer err değişkeni nil değilse (yani bir hata varsa), log.Fatalf fonksiyonu kullanılarak bir hata mesajı yazdırılır ve program sonlandırılır.
		}
		text := string(input)

		array := strings.Fields(text) // strings.Fields fonksiyonu kullanılarak, text stringi boşluklara göre bölünerek bir dizi oluşturulur. Bu dizi, array adlı bir değişkene atanır.
		for i := range array {
			array[i] = sonkontrol(array[i])
		}
		var changearray []string // changearray adında bir boş bir dizi oluşturuyoruz.
		for i, word := range array {
			if word == "(cap)" { // Eğer kelime "(cap)" ise:
				if len(changearray) > 0 { // changearray'in son elemanı büyük harfe dönüştürüyoruz.
					changearray[len(changearray)-1] = Captalized(changearray[len(changearray)-1])
				}
			} else if word == "(bin)" || word == "(bin)," { // Eğer kelime "(bin)" veya "(bin)," ise:
				if len(changearray) > 0 { // changearray'in son elemanı ikili sayıya dönüştürülür.
					changearray[len(changearray)-1] = Binary(changearray[len(changearray)-1])
				}
			} else if word == "(up)" || word == "(up)," { //  Eğer kelime "(up)" veya "(up)," ise:
				if len(changearray) > 0 { // changearray'in son elemanı büyük harfe dönüştürülür.
					changearray[len(changearray)-1] = ToUpper(changearray[len(changearray)-1])
				}
			} else if word == "(low)" || word == "(low)," { //  Eğer kelime "(low)" veya "(low)," ise:
				if len(changearray) > 0 { //  changearray'in son elemanı küçük harfe dönüştürülür.
					changearray[len(changearray)-1] = ToLower(changearray[len(changearray)-1])
				}
			} else if word == "(hex)" || word == "(hex)," { //  Eğer kelime "(hex)" veya "(hex)," ise:
				if len(changearray) > 0 { // changearray'in son elemanı onaltılık sayıya dönüştürülür.
					changearray[len(changearray)-1] = HexDecimal(changearray[len(changearray)-1])
				}
			} else if word == "a" || word == "an" || word == "A" || word == "An" { // Bir sonraki kelime ile birlikte Avovel fonksiyonuna gönderilir ve sonuç changearray'e eklenir.
				if i+1 < len(array) {
					changearray = append(changearray, Avovel(array[i+1], array[i]))
				} else {
					changearray = append(changearray, word)
				}
			} else if word == "(up," || word == "(low," || word == "(cap," { // Parantez içindeki sayıya göre, önceki kelimelerin büyük harfe dönüştürülmesi, küçük harfe dönüştürülmesi veya baş harflerinin büyük olması sağlanır.
				numstr := strings.Trim(array[i+1], ")")
				num, _ := strconv.Atoi(numstr)
				for j := 0; j < num; j++ {
					in := len(changearray) - num + j
					if in >= 0 && in < len(changearray) {
						switch word {
						case "(up,":
							changearray[in] = ToUpper(changearray[in])
						case "(low,":
							changearray[in] = ToLower(changearray[in])
						case "(cap,":
							changearray[in] = Captalized(changearray[in])
						}
					}
				}
				i++
				continue
			} else if string(word[len(word)-1]) == ")" { // Döngüye devam edilir.
				continue
			} else {
				changearray = append(changearray, word)
			}
		}
		result := strings.Join(changearray, " ")               // changearray dizisindeki kelimeler, boşluklar eklenerek birleştirilir ve result değişkenine atanır. Bu işlem, dizideki kelimelerin arasına boşluk ekleyerek bir metin oluşturur.
		result = noktalamakontrol(result)                      // noktalamakontrol fonksiyonu, metindeki noktalama işaretlerini kontrol eder. Bu adımda, olası noktalama işareti düzenlemeleri yapılır.
		result1 := sonkontrol(result)                          // sonkontrol fonksiyonu, metnin sonundaki boşlukları ve gereksiz karakterleri kontrol eder ve düzenler.
		err = os.WriteFile(os.Args[2], []byte(result1), 0o644) /* result1 metni, ikinci komut satırı argümanı olarak verilen dosyaya yazılır. Bu işlemde WriteFile fonksiyonu kullanılır.
		Dosyaya yazma işlemi sırasında bir hata oluşması durumunda, hata mesajıyla birlikte program sonlandırılır.*/
		if err != nil {
			log.Fatalf("Error writing to output file: %v", err)
		}
	} else { //  Eğer program çalıştırılırken yeterli argüman sağlanmazsa, yetersiz kullanım uyarısı verilir.
		log.Println("Usage: go run main.go <input_file> <output_file>")
	}
}

func sonkontrol(str string) string {
	str = strings.ReplaceAll(str, ".  ", ". ")
	str = strings.ReplaceAll(str, " ,", ", ")
	str = strings.ReplaceAll(str, ",  ", ", ")
	str = strings.ReplaceAll(str, " ;", ";")
	str = strings.ReplaceAll(str, " :", ":")
	str = strings.ReplaceAll(str, " !", "!")
	str = strings.ReplaceAll(str, " ?", "?")
	str = strings.ReplaceAll(str, " ...", "...")
	str = strings.ReplaceAll(str, ". . . ", "...")
	str = strings.ReplaceAll(str, "' ", "'")
	str = strings.ReplaceAll(str, ". '", ".'")
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return ' '
		}
		return r
	}, str)
}

func noktalamakontrol(s string) string {
	re := regexp.MustCompile(`\s*([,.:;!?]+)\s*`) // regexp.MustCompile fonksiyonu kullanılarak, bir düzenli ifade oluşturulur. Bu düzenli ifade, metindeki noktalama işaretlerini ve bu işaretlerin etrafındaki boşlukları kontrol etmek için kullanılır.
	return re.ReplaceAllString(s, "$1 ")          // re.ReplaceAllString fonksiyonu, oluşturulan düzenli ifadeye göre metindeki noktalama işaretlerini kontrol eder ve gereksiz boşlukları temizler. $1 ifadesi, düzenli ifadenin eşleştiği noktalama işaretlerini temsil eder.

}

func Avovel(s, bir string) string {
	vovel := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'} // vovel adında bir rune dizisi oluşturulur. Bu dizi, sesli harfleri içerir.
	isVovel := false                                                  // isVovel adında bir boolean değişkeni oluşturulur ve başlangıçta false olarak ayarlanır.
	for _, char := range vovel {                                      // range döngüsü kullanılarak, s kelimesinin her bir karakteri üzerinde dönülür.
		if rune(s[0]) == char {
			isVovel = true
			break
		}
	}
	if isVovel {
		if bir == "a" {
			bir = "an"
		} else if bir == "A" {
			bir = "An"
		}
	} else {
		if bir == "an" || bir == "An" {
			bir = "a"
		}
	}
	return bir
}

func Captalized(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(s[:1]) + s[1:] // s[:1] ifadesiyle, metnin ilk karakteri alınır.
	}
	return s
}

func Binary(s string) string {
	üs := len(s) - 1 //  üs adında bir değişken oluşturulur ve bu değişkene, s stringinin uzunluğundan 1 çıkarılarak başlangıç değeri atanır.
	sum := 0         // sum adında bir değişken oluşturulur ve başlangıçta 0 olarak ayarlanır. Bu değişken, hesaplanacak ondalık değeri tutar.
	for i := 0; i < len(s); i++ {
		eklenecek := Power(2, üs)
		if s[i] == '1' {
			sum += eklenecek
		}
		üs--
	}
	return strconv.Itoa(sum)
}

func Power(base, power int) int {
	result := 1                  //  result adında bir değişken oluşturulur ve başlangıçta 1 olarak ayarlanır. Bu değişken, hesaplanacak üssün sonucunu tutar.
	for i := 0; i < power; i++ { // i değişkeni 0'dan power değişkenine kadar artırılır.
		result *= base // result değişkeni, result ile base çarpılarak güncellenir.
	}
	return result
}

func HexDecimal(hexString string) string {
	decimal, _ := strconv.ParseInt(hexString, 16, 64)
	/* decimal adında bir değişken oluşturulur ve strconv.ParseInt(hexString, 16, 64) fonksiyonu kullanılarak onaltılı stringin ondalık değeri elde edilir. Bu fonksiyon, hexString parametresini onaltılı olarak yorumlar ve 64 bitlik bir integer olarak döndürür.*/
	return strconv.Itoa(int(decimal)) // strconv.Itoa(int(decimal)) ifadesi kullanılarak, decimal değişkeni stringe dönüştürülür ve döndürülür.
}

func ToUpper(s string) string {
	sonuc := strings.ToUpper(s) // strings.ToUpper() işlevi kullanılarak girdi stringi tamamen büyük harfe dönüştürülür.
	return sonuc
}

func ToLower(s string) string {
	sonuc := strings.ToLower(s) // strings.ToLower() işlevi kullanılarak girdi stringi tamamen küçük harfe dönüştürülür.
	return sonuc
}
