package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Client model struct.
type Client struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Rep         string    `json:"rep" db:"rep"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// ClientsReponse model struct.
type ClientsReponse struct {
	Status  int     `json:"status"`
	Data    Clients `json:"data"`
	Message string  `json:"message"`
}

// Clients array model struct of Team.
type Clients []Client

// String converts the struct into a string value.
func (c Client) String() string {
	jp, err := json.Marshal(c)
	if err != nil {
		return ""
	}

	return string(jp)
}

// Validate checks the information from given team
func (c *Client) Validate(clients Clients) *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Name, Name: "Name", Message: "Name can not be empty"},
		&validators.StringIsPresent{Field: c.PhoneNumber, Name: "PhoneNumber", Message: "PhoneNumber can not be empty"},
		&validators.StringIsPresent{Field: c.Rep, Name: "Rep", Message: "Rep can not be empty"},

		&validators.FuncValidator{
			Name:    "Name",
			Field:   c.Name,
			Message: "Client '%v' already exists.",
			Fn: func() bool {
				for _, client := range clients {
					if strings.EqualFold(c.Name, client.Name) {
						return false
					}
				}

				return true
			},
		},
	)
}
