package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciado")
	defer log.Println("Request finalizado")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processada com sucesso")
		w.Write([]byte("Request processada com sucesso"))

	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
	}
}
