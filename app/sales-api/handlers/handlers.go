// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"github.com/brabete/golang-service/business/auth"
	"github.com/brabete/golang-service/business/mid"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"

	"github.com/brabete/golang-service/foundation/web"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth, db *sqlx.DB) *web.App {

	// Construct the web.App which holds all routes as well as common Middleware.
	app := web.NewApp(shutdown, mid.Logger(log), mid.Error(log), mid.Metrics(),  mid.Panics(log))

	// Register health check endpoint. This route is not authenticated.
	c := check{
		build: build,
		db:    db,
	}

	// Register health check endpoint. This route is not authenticated.
	app.Handle(http.MethodGet, "/health", c.health)

	return app
}
