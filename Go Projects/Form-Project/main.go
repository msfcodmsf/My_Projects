package main

import (
	"form-project/allhandlers" 
	"form-project/datahandlers" // Veritabanı bağlantı bilgileri
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	datahandlers.SetDB() // Bu fonksiyon, veritabanı dosyasının yolunu ve diğer gerekli ayarları alarak bağlantıyı başlatır.
	defer datahandlers.DB.Close()

	datahandlers.CreateTables() // fonksiyonu ile veritabanında gerekli tablolar (örneğin, kullanıcı bilgileri, form verileri) oluşturulur.

	allhandlers.Allhandlers() //  fonksiyonu, HTTP isteklerini karşılayacak işleyicileri (handler) tanımlar ve kaydeder. 
	// Bu işleyiciler, /form, /submit gibi farklı URL yollarına gelen istekleri ele alır.

	log.Println("Server started at :8065")
	http.ListenAndServe(":8065", nil)
}
