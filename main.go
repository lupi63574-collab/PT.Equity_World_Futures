package main

import (
	"database/sql"
	"log"
	"net/http"
	"path/filepath"

	"go-profil-perusahaan/handler" // Pastikan modul ini sesuai dengan nama di go.mod Anda

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connectDatabase() {
	var err error

	// Koneksi ke MySQL XAMPP
	dsn := "root:@tcp(127.0.0.1:3306)/pt_equityworldfutures?parseTime=true"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal membuka koneksi database:", err)
	}

	// Cek koneksi database
	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	log.Println("Database pt_equityworldfutures berhasil terhubung!")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Menyajikan index.html dari folder view
	tmplPath := filepath.Join("view", "index.html")
	http.ServeFile(w, r, tmplPath)
}

func main() {

	// Hubungkan ke database
	connectDatabase()

	// Tutup koneksi saat program berhenti
	defer db.Close()

	// Routing
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/nasabah", handler.GetNasabahHandler(db)) // Menghubungkan route /nasabah ke handler nasabah

	log.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}