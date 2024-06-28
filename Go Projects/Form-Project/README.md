# Form-Project
* Bu proje, Go programlama dili kullanılarak geliştirilmiş basit bir forum uygulamasıdır. Kullanıcılar gönderi paylaşabilir, yorum yapabilir, gönderi ve yorumları beğenebilir/beğenmeyebilirler. Ayrıca Google ve GitHub ile OAuth tabanlı oturum açma özelliği de bulunmaktadır.

## Özellikler

* Gönderi Oluşturma: Kullanıcılar yeni gönderiler oluşturabilirler.

* Yorum Yapma: Kullanıcılar gönderilere yorum yapabilirler.

* Beğeni/Beğenmeme: Kullanıcılar gönderi ve yorumları beğenebilir veya beğenmeyebilirler.

* Oturum Açma/Kapatma: Kullanıcılar kayıt olabilir, oturum açabilir ve oturumlarını kapatabilirler.

* Google/GitHub OAuth: Google veya GitHub hesapları ile oturum açma imkanı.

* Şifre Sıfırlama: (Henüz tam olarak uygulanmamış)

## Kullanılan Teknolojiler

* Go: Programlama dili.

* SQLite3: Veritabanı.

* HTML/Templates: Görünüm katmanı için HTML şablonları.

* bcrypt: Şifreleri güvenli bir şekilde hash'lemek için.

* uuid: Benzersiz oturum tokenları oluşturmak için.

* validator: Kullanıcı girdilerini doğrulamak için.

* oauth2: OAuth 2.0 protokolünü uygulamak için.

## Kullanılan Kütüphaneler
```go
github.com/mattn/go-sqlite3: SQLite3 veritabanı sürücüsü.

github.com/go-playground/validator: Veri doğrulama için kullanılan bir kütüphane.

golang.org/x/oauth2: OAuth 2.0 istemci uygulamaları geliştirmek için kullanılan Go paketi.

golang.org/x/oauth2/google: Google OAuth 2.0 sağlayıcısı ile entegrasyon için kullanılır.

golang.org/x/oauth2/github: GitHub OAuth 2.0 sağlayıcısı ile entegrasyon için kullanılır.
```
## Proje Yapısı
```go
main.go: Uygulamanın başlangıç noktası. Veritabanı bağlantısını kurar, tabloları oluşturur ve HTTP sunucusunu başlatır.

allhandlers paketi: Tüm HTTP istek işleyicilerini (handler) içerir.

datahandlers paketi: Veritabanı işlemleri ve oturum yönetimi ile ilgili fonksiyonları içerir.

homehandlers paketi: Ana sayfa, kayıt, oturum açma, oturum kapatma ve şifre sıfırlama işlemlerini işler.

morehandlers paketi: Kullanıcı profili ve ilgili işlemleri işler.

posthandlers paketi: Gönderi oluşturma, yorum yapma, silme, oy verme ve gönderi görüntüleme işlemlerini işler.

utils paketi: Hata yönetimi gibi yardımcı fonksiyonları içerir.
```
## Kurulum
SQLite3'ü Kurun: SQLite3 veritabanını sisteminize kurun.
Gerekli Go Paketlerini İndirin: go get komutu ile gerekli Go paketlerini indirin:
Bash
```go
go get github.com/mattn/go-sqlite3
go get github.com/go-playground/validator
go get golang.org/x/oauth2
go get golang.org/x/oauth2/google
go get golang.org/x/oauth2/github
```
Kodu dikkatli kullanın.
content_copy
config.json Dosyasını Oluşturun: Google ve GitHub OAuth için gerekli istemci ID'si ve istemci sırrı bilgilerini içeren bir config.json dosyası oluşturun.
Uygulamayı Çalıştırın: go run main.go komutu ile uygulamayı çalıştırın.