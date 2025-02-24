package main

import (
	"log"
	"net/http"

	"example.com/bbb/internal/api"
	"example.com/bbb/internal/auth"
	"example.com/bbb/internal/db"
	"example.com/bbb/internal/utils"
)

func main() {

	// Load environment variables from .env
	utils.LoadEnv()

	if err := db.InitDB(); err != nil {
		log.Fatalf("Error initalizing MongoDB(in main): %v", err)
	}

	if err := auth.InitFireBase(); err != nil {
		log.Fatalf("Error initalizing Firebase: %v", err)
	}

	http.HandleFunc("/", api.HellWorld)
	http.HandleFunc("/add-business", api.AddBusiness)
	http.ListenAndServe(":8080", nil)
}
