package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Langsung menyajikan file index.html dari folder view
	tmplPath := filepath.Join("view", "index.html")
	http.ServeFile(w, r, tmplPath)
}

func main() {
	// Routing ke halaman utama
	http.HandleFunc("/", homeHandler)

	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}