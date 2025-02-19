package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/escoutdoor/kotopes/api-gateway/internal/httputil"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	"github.com/escoutdoor/kotopes/api-gateway/internal/utils/urlparam"
	"github.com/go-chi/render"
)

func (c *controller) delete(w http.ResponseWriter, r *http.Request) {
	favoriteID, err := urlparam.GetUUIDParameter(r, favoriteIDParameter)
	if err != nil {
		render.Render(w, r, httputil.ErrBadRequest(err))
		return
	}

	userID, err := httputil.ExtractUserIDFromCtx(r)
	if err != nil {
		render.Render(w, r, httputil.ErrUnauthorized(err))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	in := &model.DeleteFavorite{
		FavoriteID: favoriteID,
		UserID:     userID,
	}
	err = c.favoriteClient.Delete(ctx, in)
	if err != nil {
		httputil.HandleGrpcError(w, r, err)
		return
	}

	render.NoContent(w, r)
}
