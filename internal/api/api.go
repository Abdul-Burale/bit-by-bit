package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/bbb/internal/db"
	"example.com/bbb/internal/models"
)

func HellWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
	fmt.Println("End point")
}

func AddBusiness(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Add Business")
	fmt.Println("End point: AddBusiness")

	var businessRequest models.BusinessRequest

	// TODO: Will need to set my own marshalling here.
	// TODO: Decided in the front end.
	err := json.NewDecoder(r.Body).Decode(&businessRequest)
	if err != nil {
		log.Printf("Error decoding request body %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("Reponse: %v", r)

	// TODO: Wrap this into a function
	business := models.Business{
		Name:     businessRequest.Name,
		Category: businessRequest.Category,
		Address: models.Address{
			Street:   businessRequest.Street,
			City:     businessRequest.City,
			Postcode: businessRequest.Postcode,
			District: businessRequest.City,
		},
		Description: businessRequest.Description,
	}
	newBusiness, err := db.CreateBusiness(&business)
	if err != nil {
		http.Error(w, " end point was not succuessfull in adding the business", http.StatusInternalServerError)
		return
	}

	// Return the reponse with the newly created business
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBusiness)
}
