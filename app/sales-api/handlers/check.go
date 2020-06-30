package handlers

import (
	"context"
	"github.com/brabete/golang-service/foundation/web"
	"net/http"
)

func health(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}
