package valueobject

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Amount    int       `json:"amount" bson:"amount"`
	From      uuid.UUID `json:"from" bson:"from"`
	To        uuid.UUID `json:"to" bson:"to"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}