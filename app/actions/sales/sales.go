package sales

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

	sales := models.Sales{}
	if err := tx.All(&sales); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.SalesReponse{
		Status:  http.StatusOK,
		Data:    sales,
		Message: "Sales list",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Show(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	sale := models.Sale{}
	if err := tx.Find(&sale, c.Param("id")); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.SalesReponse{
		Status:  http.StatusOK,
		Data:    models.Sales{sale},
		Message: "Sale information",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Create(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	sale := models.Sale{}
	err := json.NewDecoder(c.Request().Body).Decode(&sale)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON("Error parsing Sale"))
	}

	if verrs := sale.Validate(); verrs.HasAny() {
		response := models.SalesReponse{
			Status:  http.StatusUnprocessableEntity,
			Data:    models.Sales{sale},
			Message: verrs.Error(),
		}

		return c.Render(http.StatusUnprocessableEntity, r.JSON(response))
	}

	if err := tx.Create(&sale); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.SalesReponse{
		Status:  http.StatusCreated,
		Data:    models.Sales{sale},
		Message: "Sale was created successfully.",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func Delete(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON("Error getting DB connection"))
	}

	sale := models.Sale{
		ID: uuid.FromStringOrNil(c.Param("id")),
	}
	if err := tx.Destroy(&sale); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	response := models.SalesReponse{
		Status:  http.StatusOK,
		Data:    models.Sales{},
		Message: "Sale was removed successfully.",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}
