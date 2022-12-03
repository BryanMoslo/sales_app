package teams

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

	teams := models.Teams{}
	if err := tx.All(&teams); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.TeamsReponse{
		Status:  http.StatusOK,
		Data:    teams,
		Message: "Teams list",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Employees(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	employees := models.Employees{}
	if err := tx.Where("team_id = ?", c.Param("id")).All(&employees); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.EmployeesReponse{
		Status:  http.StatusOK,
		Data:    employees,
		Message: "Teams employees",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Show(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	team := models.Team{}
	if err := tx.Find(&team, c.Param("id")); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.TeamsReponse{
		Status:  http.StatusOK,
		Data:    models.Teams{team},
		Message: "Team information",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Create(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	team := models.Team{}
	err := json.NewDecoder(c.Request().Body).Decode(&team)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON("Error parsing Team"))
	}

	teams := models.Teams{}
	if err := tx.All(&teams); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	if verrs := team.Validate(teams); verrs.HasAny() {
		response := models.TeamsReponse{
			Status:  http.StatusUnprocessableEntity,
			Data:    models.Teams{team},
			Message: verrs.Error(),
		}

		return c.Render(http.StatusUnprocessableEntity, r.JSON(response))
	}

	if err := tx.Create(&team); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.TeamsReponse{
		Status:  http.StatusCreated,
		Data:    models.Teams{team},
		Message: "Team was created successfully.",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Delete(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	team := models.Team{
		ID: uuid.FromStringOrNil(c.Param("id")),
	}
	if err := tx.Destroy(&team); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.TeamsReponse{
		Status:  http.StatusOK,
		Data:    models.Teams{},
		Message: "Team was removed successfully.",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}
