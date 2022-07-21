package controller

import (
	"Pratesis/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// type barang struct {
// 	iD            int64  `json:"kd_barang,pk"`
// 	name          string `json:"nm_barang"`
// 	State         string `json:"status"`
// }

type Response struct {
	// Status  int             `json:"status"`
	// Message string          `json:"message"`
	Data []models.Barang `json:"Data"`
}

// Ambil semua data barang
func AmbilSemuaBarang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// memanggil models AmbilSemuaBuku
	arrayOfBrg, err := models.AmbilSemuaBarang()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response Response
	// response.Status = 1
	// response.Message = "Success"
	response.Data = arrayOfBrg

	// kirim semua response
	json.NewEncoder(w).Encode(response)
}

// AmbilBarang mengambil single data dengan parameter id
func AmbilBarang(w http.ResponseWriter, r *http.Request) {
	// kita set headernya
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// dapatkan idbuku dari parameter request, keynya adalah "id"
	params := mux.Vars(r)

	// konversi id dari tring ke int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// memanggil models ambilsatubuku dengan parameter id yg nantinya akan mengambil single data
	brg, err := models.AmbilSatuBarang(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data Barang. %v", err)
	}

	// kirim response
	json.NewEncoder(w).Encode(brg)
}
