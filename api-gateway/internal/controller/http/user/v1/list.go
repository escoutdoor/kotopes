package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/escoutdoor/kotopes/api-gateway/internal/controller/http/user/v1/dto"
	"github.com/escoutdoor/kotopes/api-gateway/internal/converter"
	"github.com/escoutdoor/kotopes/api-gateway/internal/httputil"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	"github.com/go-chi/render"
)

func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	in := &model.ListUsers{}
	c.getListQuery(r, in)

	users, err := c.userClient.List(ctx, in)
	if err != nil {
		httputil.HandleGrpcError(w, r, err)
		return
	}

	render.JSON(w, r, dto.ListUsersResponse{Users: converter.ToDTOFromUsers(users)})
}

func (c *controller) getListQuery(r *http.Request, in *model.ListUsers) {
	userIDsQuery := r.URL.Query()["user_id"]
	if len(userIDsQuery) > 0 {
		for _, pid := range userIDsQuery {
			in.UserIDs = append(in.UserIDs, pid)
		}
	}
}
