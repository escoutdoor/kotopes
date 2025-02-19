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

func (c *controller) get(w http.ResponseWriter, r *http.Request) {
	id, err := getPetIDParameter(r)
	if err != nil {
		render.Render(w, r, httputil.ErrBadRequest(err))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	pet, err := c.petClient.Get(ctx, id)
	if err != nil {
		httputil.HandleGrpcError(w, r, err)
		return
	}

	render.JSON(w, r, dto.GetPetResponse{Pet: converter.ToDTOFromPet(pet)})
}
