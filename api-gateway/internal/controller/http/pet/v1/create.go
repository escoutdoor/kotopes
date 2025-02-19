package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/escoutdoor/kotopes/api-gateway/internal/controller/http/pet/v1/dto"
	"github.com/escoutdoor/kotopes/api-gateway/internal/converter"
	"github.com/escoutdoor/kotopes/api-gateway/internal/httputil"

	"github.com/go-chi/render"
)

func (c *controller) create(w http.ResponseWriter, r *http.Request) {
	req := dto.CreatePetRequest{}
	if err := render.DecodeJSON(r.Body, &req); err != nil {
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

	id, err := c.petClient.Create(ctx, converter.ToCreatePetFromDTO(&req, userID))
	if err != nil {
		httputil.HandleGrpcError(w, r, err)
		return
	}

	render.JSON(w, r, dto.CreatePetResponse{ID: id})
}
