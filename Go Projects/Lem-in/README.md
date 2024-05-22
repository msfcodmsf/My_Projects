# LEM-IN
> KOD AÇIKLAMASI
* Bu projenin amacı, bir dosyadan alınan verilerle bir grafik oluşturmak, tüm olası yolları bulmak ve karıncaların bu yolları en verimli şekilde kullanarak hedefe ulaşmasını sağlamaktır. Fonksiyonlar, dosya okuma, veri işleme, graf yapısı oluşturma, yolları bulma, filtreleme ve hareketleri simüle etme işlemlerini gerçekleştirmek için kullanılır.

> Projede kullandığımız kütüphaneler.
```go
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
-bufio: Dosyadan satır satır verileri okumak için kullanılır. 
-fmt: Konsola veri yazdırmak ve formatlamak için kullanılır.
-os: Dosya okuma ve yazma işlemleri için kullanılır.
-sort: Bulunan yolları uzunluklarına göre sıralamak için kullanılır.
-strconv: Dosyadan okunan verileri (örneğin karınca sayısı
ve koordinatlar) integer (tam sayı) gibi veri tiplerine dönüştürmek için kullanılır.
-strings: Dosyadan okunan satırları parçalamak, birleştirmek ve analiz etmek için kullanılır.
```

> Dosya Okuma
* İlk olarak, ReadFile fonksiyonu dosyadan verileri okur ve bunları işler.
* Dosya içeriği odalar ve bağlantılar (links) hakkında bilgiler içerir.
Başlangıç ve bitiş odaları belirlenir.
Karıncaların sayısı alınır.

> Grafik Oluşturma
* Daha sonra, CreateGraphFromRoomsAndLinks fonksiyonu odaları ve bağlantıları kullanarak bir grafik oluşturur.
* Her oda bir düğüm (node) olarak temsil edilir.
* Her bağlantı bir kenar (edge) olarak temsil edilir.

> Yol Bulma
* FindAllPaths fonksiyonu, başlangıç ve bitiş odaları arasındaki tüm olası yolları bulur.
* Derinlik öncelikli arama (DFS) kullanılarak tüm yollar bulunur.
* Bulunan yollar bir liste olarak saklanır.

> Karınca Hareketleri
* Son olarak, MoveAnts fonksiyonu karıncaların başlangıç odasından bitiş odasına nasıl hareket edeceğini simüle eder.
* Her karınca için en kısa yol belirlenir.
* Karıncaların her adımda hangi odada oldukları takip edilir.

> Kullanım Örneği
```go
go run main.go input.txt
```

>Sonuç
* Karıncaların hareketlerini ve en kısa yolları ekranda görebilirsiniz. İşte örnek bir çıktı:
```go
L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3 L4-2
L3-1 L4-3
L4-1
```



