package mid

import (
	"context"
	"github.com/brabete/golang-service/foundation/web"
	"log"
	"net/http"
	"time"
)

// Logger ...
func Logger(log *log.Logger) web.Middleware {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// If the context is missing this value, request the service
			// to be shutdown gracefully.
			v, ok := ctx.Value(web.KeyValues).(*web.Values)
			if !ok {
				return web.NewShutdownError("web value missing from context")
			}

			err := handler(ctx, w, r)

			log.Printf("%s : (%d) : %s %s -> %s (%s)",
				v.TraceID, v.StatusCode,
				r.Method, r.URL.Path,
				r.RemoteAddr, time.Since(v.Now),
			)

			// Return the error so it can be handled further up the chain.
			return err
		}
		return h
	}
	return m
}
