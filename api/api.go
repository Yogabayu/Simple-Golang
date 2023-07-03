package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// User struct represents the structure of a user
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB

func main() {
	// Membuka koneksi ke database
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/gotest")
	//root => username untuk login
	//setelah ":" kalau ada password
	//localhost => untuk domainnya
	//3306 => untuk port
	//gotest => nama database
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inisialisasi router menggunakan Gorilla Mux
	router := mux.NewRouter()

	// Menambahkan route untuk API
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	// Menjalankan server pada port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Handler untuk mendapatkan semua pengguna
func getUsers(w http.ResponseWriter, r *http.Request) {
	// Eksekusi query untuk mendapatkan semua pengguna
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Membuat slice untuk menyimpan pengguna
	var users []User

	// Iterasi melalui setiap baris hasil query
	for rows.Next() {
		var user User

		// Membaca nilai-nilai kolom dan mengisi struktur user
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatal(err)
		}

		// Menambahkan user ke slice users
		users = append(users, user)
	}

	// Mengubah slice users menjadi JSON dan mengirimkannya sebagai response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Handler untuk membuat pengguna baru
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	// Mendekode JSON request body menjadi struktur User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Eksekusi query untuk menambahkan user ke database
	_, err = db.Exec("INSERT INTO users (name) VALUES (?)", user.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"data":    user.Name,
		"success": "berhasil menambah data data",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Handler untuk mendapatkan pengguna berdasarkan ID
func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user User

	// Eksekusi query untuk mendapatkan pengguna berdasarkan ID
	err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		log.Fatal(err)
	}

	// Mengubah pengguna menjadi JSON dan mengirimkannya sebagai response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Handler untuk memperbarui pengguna berdasarkan ID
func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user User

	// Mendekode JSON request body menjadi struktur User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Eksekusi query untuk memperbarui pengguna berdasarkan ID
	_, err = db.Exec("UPDATE users SET name = ? WHERE id = ?", user.Name, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengirimkan response dengan status OK (200)
	response := map[string]interface{}{
		"id":      id,
		"success": "berhasil tambah data",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Handler untuk menghapus pengguna berdasarkan ID
func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Eksekusi query untuk menghapus pengguna berdasarkan ID
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	// Mengirimkan response dengan status OK (200)
	response := map[string]interface{}{
		"success": "berhasil menghapus data",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
