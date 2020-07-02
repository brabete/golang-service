package handlers

import (
	"context"
	"github.com/brabete/golang-service/business/auth"
	"github.com/brabete/golang-service/business/data/product"
	"github.com/brabete/golang-service/foundation/web"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"net/http"
)

type productHandlers struct {
	db *sqlx.DB
}

func (h *productHandlers) list(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	products, err := product.List(ctx, h.db)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, products, http.StatusOK)
}

func (h *productHandlers) retrieve(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	params := web.Params(r)
	prod, err := product.One(ctx, h.db, params["id"])
	if err != nil {
		switch err {
		case product.ErrInvalidID:
			return web.NewRequestError(err, http.StatusBadRequest)
		case product.ErrNotFound:
			return web.NewRequestError(err, http.StatusNotFound)
		default:
			return errors.Wrapf(err, "ID: %s", params["id"])
		}
	}

	return web.Respond(ctx, w, prod, http.StatusOK)
}

func (h *productHandlers) create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	claims, ok := ctx.Value(auth.Key).(auth.Claims)
	if !ok {
		return web.NewShutdownError("claims missing from context")
	}

	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	var np product.NewProduct
	if err := web.Decode(r, &np); err != nil {
		return errors.Wrap(err, "decoding new product")
	}

	prod, err := product.Create(ctx, h.db, claims, np, v.Now)
	if err != nil {
		return errors.Wrapf(err, "creating new product: %+v", np)
	}

	return web.Respond(ctx, w, prod, http.StatusCreated)
}
