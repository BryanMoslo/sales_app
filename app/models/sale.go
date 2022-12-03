package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Sale model struct.
type Sale struct {
	ID        uuid.UUID `json:"id" db:"id"`
	OfferID   uuid.UUID `json:"offer_id" db:"offer_id"`
	TeamID    uuid.UUID `json:"team_id" db:"team_id"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// SalesReponse model struct.
type SalesReponse struct {
	Status  int    `json:"status"`
	Data    Sales  `json:"data"`
	Message string `json:"message"`
}

// Sales array model struct of Sale.
type Sales []Sale

// String converts the struct into a string value.
func (s Sale) String() string {
	jp, err := json.Marshal(s)
	if err != nil {
		return ""
	}

	return string(jp)
}

// Validate checks the information from given team
func (s *Sale) Validate() *validate.Errors {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: s.OfferID, Name: "Offer", Message: "Offer can not be empty"},
		&validators.UUIDIsPresent{Field: s.TeamID, Name: "TeamID", Message: "Team can not be empty"},
	)
}
