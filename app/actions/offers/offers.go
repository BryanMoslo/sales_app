package offers

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

	offers := models.Offers{}
	if err := tx.All(&offers); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.OffersReponse{
		Status:  http.StatusOK,
		Data:    offers,
		Message: "Offers list",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Show(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	offer := models.Offer{}
	if err := tx.Find(&offer, c.Param("id")); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.OffersReponse{
		Status:  http.StatusOK,
		Data:    models.Offers{offer},
		Message: "Offer information",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Create(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	offer := models.Offer{}
	err := json.NewDecoder(c.Request().Body).Decode(&offer)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON("Error parsing Offer"))
	}

	if verrs := offer.Validate(); verrs.HasAny() {
		response := models.OffersReponse{
			Status:  http.StatusUnprocessableEntity,
			Data:    models.Offers{offer},
			Message: verrs.Error(),
		}

		return c.Render(http.StatusUnprocessableEntity, r.JSON(response))
	}

	if err := tx.Create(&offer); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.OffersReponse{
		Status:  http.StatusCreated,
		Data:    models.Offers{offer},
		Message: "Offer was created successfully.",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Delete(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	offer := models.Offer{
		ID: uuid.FromStringOrNil(c.Param("id")),
	}
	if err := tx.Destroy(&offer); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.OffersReponse{
		Status:  http.StatusOK,
		Data:    models.Offers{},
		Message: "Offer was removed successfully.",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}
