package main

import (
	"fmt"
	"go-crud-postgres/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	// Use the router from the imported package
	r := routes.Router()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}