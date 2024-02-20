# ASCII-ART
* Bu projemizde kullanıcıdan alınan bir metni ASCII sanatına dönüştürür.
Kod, metni kelimelerine ayırır ve her kelimeyi ASCII karakterlerinden oluşan bir diziye dönüştürür.
Daha sonra bu diziyi 8x8'lik bir tabloda gösterir.

> Projede kullanılan kütüphaneler
```go
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
os: 	 Dosya okuma ve yazma işlemleri için kullanılır.
strings: Metin işlemleri için kullanılır.
bufio:   Bir dosyadaki veya standart girişten veri okumak/veriyi yazmak için kullanılır.
fmt:     Formatlanmış çıktıyı görüntülmek için kullanılır.
*/
```

> Kullanım
* Programı çalıştırmak için aşağıdaki adımları izleyin:

* Go programlama dilini bilgisayarınıza kurun.
* Programın kodunu bir ".go" uzantılı dosyaya kaydedin.
* Komut satırında aşağıdaki komutu çalıştırın:
```go
go run <dosya_adı>.go <metin>
```
> Örnek:
```go
go run main.go "Hello" | cat -e
```
> Çıktı:
```go
 _    _            _    _            $
| |  | |          | |  | |           $
| |__| |    ___   | |  | |    ___    $
|  __  |   / _ \  | |  | |   / _ \   $
| |  | |  |  __/  | |  | |  | (_) |  $
|_|  |_|   \___|  |_|  |_|   \___/   $
                                     $
                                     $
```

> Örnek:
```go
go run main.go "Hello\nThere" | cat -e
```
> Çıktı:
```go
 _    _            _    _            $
| |  | |          | |  | |           $
| |__| |    ___   | |  | |    ___    $
|  __  |   / _ \  | |  | |   / _ \   $
| |  | |  |  __/  | |  | |  | (_) |  $
|_|  |_|   \___|  |_|  |_|   \___/   $
                                     $
                                     $
 _______    _                               $
|__   __|  | |                              $
   | |     | |__      ___    _ __     ___   $
   | |     |  _ \    / _ \  | '__|   / _ \  $
   | |     | | | |  |  __/  | |     |  __/  $
   |_|     |_| |_|   \___|  |_|      \___|  $
                                            $
                                            $
```
