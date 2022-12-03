package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

var (
	SellerRole = "seller"
	TypistRole = "typist"
	LegalRole  = "legal"
	LeaderRole = "leader"
)

// Employee model struct.
type Employee struct {
	ID        uuid.UUID `json:"id" db:"id"`
	TeamID    uuid.UUID `json:"team_id" db:"team_id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Role      string    `json:"role" db:"role"`
	Rate      float64   `json:"rate" db:"rate"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// EmployeesReponse model struct.
type EmployeesReponse struct {
	Status  int       `json:"status"`
	Data    Employees `json:"data"`
	Message string    `json:"message"`
}

// Employees array model struct of Employee.
type Employees []Employee

// String converts the struct into a string value.
func (e Employee) String() string {
	jp, err := json.Marshal(e)
	if err != nil {
		return ""
	}

	return string(jp)
}

// Validate checks the information from given team
func (e *Employee) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.FirstName, Name: "FirstName", Message: "First Name can not be empty"},
		&validators.StringIsPresent{Field: e.LastName, Name: "LastName", Message: "Last Name can not be empty"},
		&validators.StringIsPresent{Field: e.Role, Name: "Role", Message: "Role can not be empty"},

		&validators.FuncValidator{
			Name:    "Role",
			Field:   e.Role,
			Message: "Role '%v' is not part of our roles.",
			Fn: func() bool {
				return e.Role == SellerRole || e.Role == TypistRole || e.Role == LegalRole || e.Role == LeaderRole
			},
		},
	)
}
