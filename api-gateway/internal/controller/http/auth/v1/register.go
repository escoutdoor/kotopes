package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/escoutdoor/kotopes/api-gateway/internal/controller/http/auth/v1/dto"
	"github.com/escoutdoor/kotopes/api-gateway/internal/converter"
	"github.com/escoutdoor/kotopes/api-gateway/internal/httputil"
	"github.com/go-chi/render"
)

func (c *controller) register(w http.ResponseWriter, r *http.Request) {
	req := dto.RegisterRequest{}
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Render(w, r, httputil.ErrBadRequest(err))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	id, err := c.authClient.Register(ctx, converter.ToRegisterFromDTO(&req))
	if err != nil {
		httputil.HandleGrpcError(w, r, err)
		return
	}

	render.JSON(w, r, dto.RegisterResponse{ID: id})
}
