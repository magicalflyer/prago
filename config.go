package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Fungsi koneksi ke database
func CreateConnection() *sql.DB {

	// Membuka setting dbase melalui file .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Tidak bisa load file .env")
	}

	// Membuka koneksi ke dbase
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	// Cek sambungan koneksi
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Terhubung ke dbase")

	return db
}

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}
