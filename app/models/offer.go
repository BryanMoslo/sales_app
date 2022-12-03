package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Offer model struct.
type Offer struct {
	ID          uuid.UUID `json:"id" db:"id"`
	ClientID    uuid.UUID `json:"client_id" db:"client_id"`
	Description string    `json:"description" db:"description"`
	Industry    string    `json:"industry" db:"industry"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// OffersReponse model struct.
type OffersReponse struct {
	Status  int    `json:"status"`
	Data    Offers `json:"data"`
	Message string `json:"message"`
}

// Offers array model struct of Offer.
type Offers []Offer

// String converts the struct into a string value.
func (o Offer) String() string {
	jp, err := json.Marshal(o)
	if err != nil {
		return ""
	}

	return string(jp)
}

// Validate checks the information from given team
func (o *Offer) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: o.Description, Name: "Description", Message: "Name can not be empty"},
		&validators.StringIsPresent{Field: o.Industry, Name: "Industry", Message: "Industry can not be empty"},
		&validators.UUIDIsPresent{Field: o.ClientID, Name: "Client", Message: "Client can not be empty"},

		&validators.FuncValidator{
			Name:    "Name",
			Field:   o.Industry,
			Message: "Industry '%v' is not part of our services.",
			Fn: func() bool {
				return o.Industry == HealthIndustry || o.Industry == InsuranceIndustry || o.Industry == EntertainmentIndustry
			},
		},
	)
}
