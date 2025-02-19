package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/escoutdoor/kotopes/api-gateway/internal/httputil"
	"github.com/go-chi/render"
)

func (c *controller) delete(w http.ResponseWriter, r *http.Request) {
	id, err := httputil.ExtractUserIDFromCtx(r)
	if err != nil {
		render.Render(w, r, httputil.ErrUnauthorized(err))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	err = c.userClient.Delete(ctx, id)
	if err != nil {
		httputil.HandleGrpcError(w, r, err)
		return
	}

	render.NoContent(w, r)
}
