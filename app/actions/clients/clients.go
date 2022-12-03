package clients

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

	clients := models.Clients{}
	if err := tx.All(&clients); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.ClientsReponse{
		Status:  http.StatusOK,
		Data:    clients,
		Message: "Clients list",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Offers(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	offers := models.Offers{}
	if err := tx.Where("client_id = ?", c.Param("id")).All(&offers); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.OffersReponse{
		Status:  http.StatusOK,
		Data:    offers,
		Message: "Clients offers",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Show(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	client := models.Client{}
	if err := tx.Find(&client, c.Param("id")); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.ClientsReponse{
		Status:  http.StatusOK,
		Data:    models.Clients{client},
		Message: "Client information",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Create(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	client := models.Client{}
	err := json.NewDecoder(c.Request().Body).Decode(&client)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON("Error parsing Client"))
	}

	clients := models.Clients{}
	if err := tx.All(&clients); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err.Error()))
	}

	if verrs := client.Validate(clients); verrs.HasAny() {
		response := models.ClientsReponse{
			Status:  http.StatusUnprocessableEntity,
			Data:    models.Clients{client},
			Message: verrs.Error(),
		}

		return c.Render(http.StatusUnprocessableEntity, r.JSON(response))
	}

	if err := tx.Create(&client); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err.Error()))
	}

	response := models.ClientsReponse{
		Status:  http.StatusCreated,
		Data:    models.Clients{client},
		Message: "Client was created successfully.",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Delete(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	client := models.Client{
		ID: uuid.FromStringOrNil(c.Param("id")),
	}
	if err := tx.Destroy(&client); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.ClientsReponse{
		Status:  http.StatusOK,
		Data:    models.Clients{},
		Message: "Client was removed successfully.",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}
