package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
	} else {
		for _, ch := range os.Args[1] {
			if ch >= 'a' && ch <= 'z' {
				ch = 'z' - ch + 'a'
			} else if ch >= 'A' && ch <= 'Z' {
				ch = 'Z' - ch + 'A'
			}
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}

/*
 Bu döngü, verilen argümanın her karakterini döner. Eğer karakter küçük harf ise ('a' ile 'z' arasında)
 karakteri 'z' - ch + 'a' formülüyle ters çevirir. Bu formül alfabenin bir harfini diğer uca çevirir. 
 Aynı mantık büyük harfler için de geçerlidir ('A' ile 'Z' arasında).
*/