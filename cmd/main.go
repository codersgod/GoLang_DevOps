package main

import (
	"log"
	"net/http"

	"github.com/user/cost-optimizer/app"
)

// main handles local testing execution instances.
func main() {
	log.Println("Server running locally on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", app.LocalHandler()))
}
