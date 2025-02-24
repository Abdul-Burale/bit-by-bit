package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Street   string
	City     string
	District string
	Postcode string
}

type Business struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Category    string             `bson:"category"`
	Address     Address            `bson:"inline"`
	Description string             `bson:"description"` // TODO: Find out size of string in goLang
}

type BusinessRequest struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	City        string `json:"city"`
	Street      string `json:"street"`
	Postcode    string `json:"postcode"`
	Address     string `json:"address"`
	Description string `json:"description"`
}
