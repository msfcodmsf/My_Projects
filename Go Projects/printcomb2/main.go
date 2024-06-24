package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for a := '0'; a <= '9'; a++ { // Birinci sayıyı temsil eden rakam
		for b := '0'; b <= '9'; b++ { // İkinci sayıyı temsil eden rakam
			d := b + 1
			for c := a; c <= '9'; c++ { // Üçüncü sayıyı temsil eden rakam
				for ; d <= '9'; d++ { // Dördüncü sayıyı temsil eden rakam
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(d)
					// Sayıları arasına virgül ve boşluk eklemek için kontrol sağlandı
					if a < '9' || b < '8' || c < '9' || d < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				d = '0' // Her seferinde '0' değerine sıfırlanması sağlandı
			}
		}
	}
	z01.PrintRune('\n') // Yeni satır ekleme
}
