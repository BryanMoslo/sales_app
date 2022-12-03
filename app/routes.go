package app

import (
	"net/http"

	"sales_app/app/actions/clients"
	"sales_app/app/actions/employees"
	"sales_app/app/actions/offers"
	"sales_app/app/actions/sales"
	"sales_app/app/actions/teams"
	"sales_app/app/middleware"
	"sales_app/public"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/teams", teams.List)
	root.GET("/teams/{id}", teams.Show)
	root.GET("/teams/{id}/employees", teams.Employees)
	root.POST("/teams/create", teams.Create)
	root.DELETE("/teams/{id}", teams.Delete)

	root.GET("/clients", clients.List)
	root.GET("/clients/{id}", clients.Show)
	root.GET("/clients/{id}/offers", clients.Offers)
	root.POST("/clients/create", clients.Create)
	root.DELETE("/clients/{id}", clients.Delete)

	root.GET("/offers", offers.List)
	root.GET("/offers/{id}", offers.Show)
	root.POST("/offers/create", offers.Create)
	root.DELETE("/offers/{id}", offers.Delete)

	root.GET("/sales", sales.List)
	root.GET("/sales/{id}", sales.Show)
	root.POST("/sales/create", sales.Create)
	root.DELETE("/sales/{id}", sales.Delete)

	root.GET("/employees", employees.List)
	root.GET("/employees/{id}", employees.Show)
	root.POST("/employees/create", employees.Create)
	root.DELETE("/employees/{id}", employees.Delete)

	root.ServeFiles("/", http.FS(public.FS()))
}
