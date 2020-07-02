package handlers

import (
	"context"
	"github.com/brabete/golang-service/foundation/database"
	"github.com/brabete/golang-service/foundation/web"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type check struct {
	build string
	db    *sqlx.DB
}

// If the database is not ready we will tell the client and use a 500
// status. Do not respond by just returning an error because further up in
// the call stack will interpret that as an unhandled error.
func (c *check) health(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	status := "OK"
	statusCode := http.StatusOK

	if err := database.StatusCheck(ctx, c.db); err != nil {
		status = "DB not ready"
		statusCode = http.StatusInternalServerError
	}

	health := struct {
		Version string `json:"version"`
		Status  string `json:"status"`
	}{
		Version: c.build,
		Status:  status,
	}

	return web.Respond(ctx, w, health, statusCode)
}
