package models

import (
	"Pratesis/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

// Barang schema dari tabel Barang
// kita coba dengan jika datanya null
// jika return datanya ada yg null, silahkan pake NullString, contohnya dibawah
// Penulis       config.NullString `json:"penulis"`

type Barang struct {
	Id    int64  `json:"kd_barang,pk"`
	Name  string `json:"nm_barang"`
	State string `json:"status"`
}

// ambil satu buku
func AmbilSemuaBarang() ([]Barang, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var arrayOfBrg []Barang

	// kita buat select query
	sqlStatement := `SELECT * FROM mbarang`

	// mengeksekusi sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// kita tutup eksekusi proses sql qeurynya
	defer rows.Close()

	// kita iterasi mengambil datanya
	for rows.Next() {
		var brgFromDb Barang

		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&brgFromDb.Id, &brgFromDb.Name, &brgFromDb.State)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		// masukkan kedalam slice bukus
		arrayOfBrg = append(arrayOfBrg, brgFromDb)

	}

	// return empty buku atau jika error
	return arrayOfBrg, err
}

// mengambil satu barang
func AmbilSatuBarang(id int64) (Barang, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var brgFromDb Barang

	// buat sql query
	sqlStatement := `SELECT * FROM buku WHERE id=$1`

	// eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&brgFromDb.Id, &brgFromDb.Name, &brgFromDb.State)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return brgFromDb, nil
	case nil:
		return brgFromDb, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return brgFromDb, err
}
