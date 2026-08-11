package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Menentukan handler untuk URL rute utama ("/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// WriteString mengirimkan teks langsung ke browser
		fmt.Fprintln(w, "Hello, World!")
	})

	log.Println("Server running on http://localhost:8080")
	
	// Jalankan server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}