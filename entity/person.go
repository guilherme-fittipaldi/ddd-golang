package entity

import (
	"github.com/google/uuid"
)

type Person struct {
	ID   uuid.UUID `json:"id" bson:"id"`
	Name string    `json:"name" bson:"name"`
	Age  int       `json:"age" bson:"age"`
}
