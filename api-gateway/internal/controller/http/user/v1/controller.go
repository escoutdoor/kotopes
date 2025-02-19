package v1

import (
	"net/http"

	grpc_client "github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc"
	"github.com/escoutdoor/kotopes/api-gateway/internal/utils/urlparam"
	"github.com/go-chi/chi/v5"
)

const (
	userIDParameter = "id"
)

type controller struct {
	userClient grpc_client.UserClient
}

func newController(userClient grpc_client.UserClient) *controller {
	return &controller{
		userClient: userClient,
	}
}

func Register(
	mux chi.Router,
	userClient grpc_client.UserClient,
	authMiddleware func(http.Handler) http.Handler,
) {
	r := chi.NewRouter()
	ctl := newController(userClient)

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)

		r.Patch("/", ctl.update)
		r.Delete("/", ctl.delete)
	})

	r.Get("/{id}", ctl.get)
	r.Get("/", ctl.list)

	mux.Mount("/users/v1", r)
}

func getUserIDParameter(r *http.Request) (string, error) {
	id, err := urlparam.GetUUIDParameter(r, userIDParameter)
	if err != nil {
		return "", err
	}

	return id, nil
}
