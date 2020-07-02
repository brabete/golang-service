// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"github.com/brabete/golang-service/business/auth"
	"github.com/brabete/golang-service/business/mid"
	"log"
	"net/http"
	"os"

	"github.com/brabete/golang-service/foundation/web"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Error(log), mid.Metrics(),  mid.Panics(log))

	app.Handle(http.MethodGet, "/health", health)

	return app
}
