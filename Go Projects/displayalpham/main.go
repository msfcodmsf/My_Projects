package main

import "github.com/01-edu/z01"

func main() { // tersten veya düzden, büyüklü küçüklü yazdırmanın yolu ASCII yi kullanarak +32 veya -32 yapmak.
	count := 0
	for i := 'Z'; i >= 'A'; i-- { // tersten saydırmalarda i-- olarak azalarak gittiğimize dikkat!!
		if count%2 == 0 { // count sayacının mod2 sini alıyor buda şu demek eğer çif ise zaten 0 olucak ama tekli sayılar geldiğinde sonuç 1 oluyor.
			z01.PrintRune(i)
			count++ // bir sonraki harf küçük çıksın diye countu artırıyoruz
		} else {
			z01.PrintRune(i + 32) // count tek sayıysa, küçük harfi yazdırmak için büyük harfin ASCII değerine 32 ekleyerek küçültüyoruz
			count++
		}
	}
	z01.PrintRune('\n')
}