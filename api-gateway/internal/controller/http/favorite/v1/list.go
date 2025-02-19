package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/escoutdoor/kotopes/api-gateway/internal/controller/http/favorite/v1/dto"
	"github.com/escoutdoor/kotopes/api-gateway/internal/converter"
	"github.com/escoutdoor/kotopes/api-gateway/internal/httputil"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	"github.com/go-chi/render"
)

func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	userID, err := httputil.ExtractUserIDFromCtx(r)
	if err != nil {
		render.Render(w, r, httputil.ErrUnauthorized(err))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	in := &model.ListFavorites{UserID: userID}
	c.getListQuery(r, in)

	favoriteList, err := c.favoriteClient.ListFavorites(ctx, in)
	if err != nil {
		httputil.HandleGrpcError(w, r, err)
		return
	}

	render.JSON(w, r, dto.ListFavoritesResponse{
		Favorites: converter.ToDTOFromFavoritePets(favoriteList.Favorites),
		Total:     favoriteList.Total,
	})
}

func (c *controller) getListQuery(r *http.Request, in *model.ListFavorites) {
	limitQuery := r.URL.Query().Get("limit")
	if limitQuery != "" {
		limit, _ := strconv.Atoi(limitQuery)
		in.Limit = int32(limit)
	}

	offsetQuery := r.URL.Query().Get("offset")
	if offsetQuery != "" {
		offset, _ := strconv.Atoi(offsetQuery)
		in.Offset = int32(offset)
	}
}
