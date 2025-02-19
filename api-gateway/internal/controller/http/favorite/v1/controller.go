package v1

import (
	"net/http"

	"github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc"
	"github.com/go-chi/chi/v5"
)

const (
	petIDParameter      = "id"
	favoriteIDParameter = "id"
)

type controller struct {
	favoriteClient grpc.FavoriteClient
}

func newController(favoriteClient grpc.FavoriteClient) *controller {
	return &controller{
		favoriteClient: favoriteClient,
	}
}

func Register(
	mux *chi.Mux,
	favoriteClient grpc.FavoriteClient,
	authMiddleware func(http.Handler) http.Handler,
) {
	ctl := newController(favoriteClient)
	r := chi.NewRouter()
	r.Use(authMiddleware)

	r.Post("/{id}", ctl.create)
	r.Get("/", ctl.list)
	r.Delete("/{id}", ctl.delete)

	mux.Mount("/favorites/v1", r)
}
