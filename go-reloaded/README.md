# Go-Reloaded

* Bu projemizde dosyaları basit bir araç oluşturarak noktalama işaretleri, boşluk düzeltme, kelimelerin kontrolü gibi işlemler yaptırılmıştır.
> Projede kullandığım kütüphaneler.
```go
import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

/*
os: Dosya okuma ve yazma işlemleri için kullanılır.
log: Hata ve bilgi mesajlarını yazdırmak için kullanılır.
regexp: Düzenli ifadelerle metni işlemek için kullanılır.
strconv: String ve sayısal değerler arasında dönüştürmeler yapmak için kullanılır.
strings: Metin işlemleri için kullanılır.
unicode: Unicode karakterleri ile ilgili işlemler için kullanılır.
*/
```
>  Kodun nasıl çalıştığı hakkında bir açıklama.

### `main()` Fonksiyonu
Programın giriş noktası, program başlatıldığında çalıştırılan ilk fonksiyondur. Bu durumda, `main()` içinde `readFile()`, `processData
Programı başlatır, komut satırına gelen parametreleri alır ve programı çalıştırır.

> Kodun amacı:
* Bu kod, bir metin dosyasını okuyup işleyerek kelimeler üzerinde çeşitli işlemler yapar ve işlenmiş metni yeni bir dosyaya yazar. Kodun temel işlevleri şunlardır:

> Özel Kelime İşleme:

* (cap): Kelimenin ilk harfini büyük harfe dönüştürür.
* (bin): Kelimenin ikili sayı sistemindeki karşılığını ondalık sayıya dönüştürür.
* (up): Kelimenin tamamını büyük harfe dönüştürür.
* (low): Kelimenin tamamını küçük harfe dönüştürür.
* (hex): Kelimenin onaltılık sayı sistemindeki karşılığını ondalık sayıya dönüştürür.

> Noktalama İşareti ve Boşluk Düzeltilmesi:

* Metindeki noktalama işaretleri ve boşluklar kontrol edilir ve düzeltilir.
* Fazladan boşluklar kaldırılır.
* Noktalama işaretlerinden sonra boşluk eklenir.
"a" ve "an" Kelimelerinin Kontrolü:

* "a" veya "an" kelimeleri, bir sonraki kelimenin ilk harfine göre "an" veya "a" olarak değiştirilir.
* İşlemden sonra metin, orijinal dosyanın adıyla aynı ada sahip ancak ".out" uzantılı yeni bir dosyaya yazılır.

> Kullanım Örneği
```go
go run main.go input.txt output.txt input.txt
```

* Bu metin bir örnektir. (cap)Merhaba(low) dünya! (bin)1001(up) (hex)A

> output.txt:

* Merhaba dünya! 9 A

> Kodun Özellikleri
* Kolay kullanılabilir ve anlaşılır.
* Genişletilebilir ve özelleştirilebilir.
* Farklı metin işleme görevleri için kullanılabilir.