package employees

import (
	"encoding/json"
	"net/http"
	"sales_app/app/models"
	"sales_app/app/render"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

var (
	// r is a buffalo/render Engine that will be used by actions
	// on this package to render render HTML or any other formats.
	r = render.Engine
)

func List(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	employees := models.Employees{}
	if err := tx.All(&employees); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.EmployeesReponse{
		Status:  http.StatusOK,
		Data:    employees,
		Message: "Employees list",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Show(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	employee := models.Employee{}
	if err := tx.Find(&employee, c.Param("id")); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.EmployeesReponse{
		Status:  http.StatusOK,
		Data:    models.Employees{employee},
		Message: "Employee information",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Create(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	employee := models.Employee{}
	err := json.NewDecoder(c.Request().Body).Decode(&employee)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON("Error parsing Employee"))
	}

	if verrs := employee.Validate(); verrs.HasAny() {
		response := models.EmployeesReponse{
			Status:  http.StatusUnprocessableEntity,
			Data:    models.Employees{employee},
			Message: verrs.Error(),
		}

		return c.Render(http.StatusUnprocessableEntity, r.JSON(response))
	}

	if err := tx.Create(&employee); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.EmployeesReponse{
		Status:  http.StatusCreated,
		Data:    models.Employees{employee},
		Message: "Employee was created successfully.",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Delete(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	employee := models.Employee{
		ID: uuid.FromStringOrNil(c.Param("id")),
	}
	if err := tx.Destroy(&employee); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.EmployeesReponse{
		Status:  http.StatusOK,
		Data:    models.Employees{},
		Message: "Employee was removed successfully.",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}
