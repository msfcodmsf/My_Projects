package utils

import (
	"log"
	"net/http"
)

func HandleErr(w http.ResponseWriter, err error, userMessage string, code int) {
	log.Printf("Error: %v", err)
	http.Error(w, userMessage, code)
}
