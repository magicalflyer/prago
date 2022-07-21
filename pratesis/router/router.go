package router

import (
	"Pratesis/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/barang", controller.AmbilSemuaBarang).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/barang/{id}", controller.AmbilBarang).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/barang", controller.TmbhBuku).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/barang/{id}", controller.UpdateBuku).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/barang/{id}", controller.HapusBuku).Methods("DELETE", "OPTIONS")

	return router
}
