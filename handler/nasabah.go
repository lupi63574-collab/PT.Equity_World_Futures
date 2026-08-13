package handler

import (
	"database/sql"
	"html/template"
	"net/http"
)

type Nasabah struct {
	ID    int
	Nama  string
	Email string
	Telp  string
}

func GetNasabahHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Mengambil seluruh kolom dari tabel nasabah
		rows, err := db.Query("SELECT * FROM nasabah")
		if err != nil {
			http.Error(w, "Gagal mengambil data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var listNasabah []Nasabah

		// Cek daftar nama kolom yang ada di database secara dinamis
		cols, err := rows.Columns()
		if err == nil {
			for rows.Next() {
				// Menyiapkan slice penampung dinamis berdasarkan jumlah kolom
				columns := make([]interface{}, len(cols))
				columnPointers := make([]interface{}, len(cols))
				for i := range columns {
					columnPointers[i] = &columns[i]
				}

				if err := rows.Scan(columnPointers...); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Memasukkan data ke struct Nasabah secara fleksibel
				var n Nasabah
				for i, colName := range cols {
					val := columns[i]
					var strVal string
					if val != nil {
						strVal = string(val.([]byte))
					}

					switch colName {
					case "id", "id_nasabah":
						// jika ID angka
						if intVal, ok := val.(int64); ok {
							n.ID = int(intVal)
						}
					case "nama", "nama_nasabah", "nama_lengkap":
						n.Nama = strVal
					case "email":
						n.Email = strVal
					case "telp", "no_telp", "no_telepon", "telepon":
						n.Telp = strVal
					}
				}
				listNasabah = append(listNasabah, n)
			}
		}

		tmpl, err := template.ParseFiles("view/nasabah.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, listNasabah)
	}
}