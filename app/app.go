package app

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/rs/cors"
)

var (
	root *buffalo.App
)

// App creates a new application with default settings and reading
// GO_ENV. It calls setRoutes to setup the routes for the app that's being
// created before returning it
func New() *buffalo.App {
	if root != nil {
		return root
	}

	root = buffalo.New(buffalo.Options{
		Env:         envy.Get("GO_ENV", "development"),
		SessionName: "_sales_app_session",
		PreWares: []buffalo.PreWare{cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{
				"Content-Type", "application/json",
				"Authorization",
			},
			AllowCredentials: true,
			Debug:            true,
		}).Handler},
	})

	// Setting the routes for the app
	setRoutes(root)

	return root
}
