package routes

import (
	"go-crud-postgres/controllers" // Update the import path to match your project structure

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/buku", controller.AmbilSemuaBuku).Methods("GET", "OPTIONS") // Update the controller reference
	router.HandleFunc("/api/buku/{id}", controller.AmbilBuku).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/buku", controller.TmbhBuku).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/buku/{id}", controller.UpdateBuku).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/buku/{id}", controller.HapusBuku).Methods("DELETE", "OPTIONS")

	return router
}
