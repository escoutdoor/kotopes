package v1

import (
	grpc_client "github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc"
	"github.com/go-chi/chi/v5"
)

type controller struct {
	authClient grpc_client.AuthClient
}

func newController(
	authClient grpc_client.AuthClient,
) *controller {
	return &controller{
		authClient: authClient,
	}
}

func Register(
	mux chi.Router,
	authClient grpc_client.AuthClient,
) {
	var (
		r   = chi.NewRouter()
		ctl = newController(authClient)
	)

	r.Post("/login", ctl.login)
	r.Post("/register", ctl.register)
	r.Post("/refresh", ctl.refresh)

	mux.Mount("/auth/v1", r)
}
