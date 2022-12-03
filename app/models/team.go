package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

var (
	HealthIndustry        = "health"
	InsuranceIndustry     = "insurance"
	EntertainmentIndustry = "entertainment"
)

// Team model struct.
type Team struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Industry  string    `json:"industry" db:"industry"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// TeamsReponse model struct.
type TeamsReponse struct {
	Status  int    `json:"status"`
	Data    Teams  `json:"data"`
	Message string `json:"message"`
}

// Teams array model struct of Team.
type Teams []Team

// String converts the struct into a string value.
func (t Team) String() string {
	jp, err := json.Marshal(t)
	if err != nil {
		return ""
	}

	return string(jp)
}

// Validate checks the information from given team
func (t *Team) Validate(teams Teams) *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Name, Name: "Name", Message: "Name can not be empty"},
		&validators.StringIsPresent{Field: t.Industry, Name: "Industry", Message: "Industry can not be empty"},

		&validators.FuncValidator{
			Name:    "Name",
			Field:   t.Name,
			Message: "Team '%v' already exists.",
			Fn: func() bool {
				for _, team := range teams {
					if strings.EqualFold(t.Name, team.Name) {
						return false
					}
				}

				return true
			},
		},

		&validators.FuncValidator{
			Name:    "Name",
			Field:   t.Industry,
			Message: "Industry '%v' is not part of our services.",
			Fn: func() bool {
				return t.Industry == HealthIndustry || t.Industry == InsuranceIndustry || t.Industry == EntertainmentIndustry
			},
		},
	)
}
