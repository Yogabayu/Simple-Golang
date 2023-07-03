package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Konfigurasi koneksi database MySQL
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/gotest")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Mengecek koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Terhubung ke database MySQL!")

	// Contoh penggunaan query SELECT
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Membaca hasil query
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Contoh penggunaan query INSERT
	result, err := db.Exec("INSERT INTO users(name) VALUES(?)", "John")
	if err != nil {
		log.Fatal(err)
	}

	affectedRows, _ := result.RowsAffected()
	lastInsertID, _ := result.LastInsertId()
	fmt.Printf("Rows affected: %d, Last insert ID: %d\n", affectedRows, lastInsertID)
}
