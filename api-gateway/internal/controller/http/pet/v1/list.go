package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/escoutdoor/kotopes/api-gateway/internal/controller/http/pet/v1/dto"
	"github.com/escoutdoor/kotopes/api-gateway/internal/converter"
	"github.com/escoutdoor/kotopes/api-gateway/internal/httputil"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	"github.com/go-chi/render"
)

func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	in := &model.ListPets{}
	c.getListQuery(r, in)

	pets, err := c.petClient.ListPets(ctx, in)
	if err != nil {
		httputil.HandleGrpcError(w, r, err)
		return
	}

	render.JSON(w, r, dto.ListPetsResponse{Pets: converter.ToDTOFromPets(pets)})
}

func (c *controller) getListQuery(r *http.Request, in *model.ListPets) {
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

	petIDsQuery := r.URL.Query()["pet_id"]
	if len(petIDsQuery) > 0 {
		for _, pid := range petIDsQuery {
			in.PetIDs = append(in.PetIDs, pid)
		}
	}
}
