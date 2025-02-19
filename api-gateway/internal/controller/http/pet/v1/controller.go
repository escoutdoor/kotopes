package v1

import (
	"net/http"

	grpc_client "github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc"
	"github.com/escoutdoor/kotopes/api-gateway/internal/utils/urlparam"
	"github.com/go-chi/chi/v5"
)

const (
	petIDParameter = "id"
)

type controller struct {
	petClient grpc_client.PetClient
}

func newController(petClient grpc_client.PetClient) *controller {
	return &controller{
		petClient: petClient,
	}
}

func Register(
	mux chi.Router,
	petClient grpc_client.PetClient,
	authMiddleware func(http.Handler) http.Handler,
) {
	var (
		r   = chi.NewRouter()
		ctl = newController(petClient)
	)

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)

		r.Post("/", ctl.create)
		r.Patch("/{id}", ctl.update)
		r.Delete("/{id}", ctl.delete)
	})

	r.Get("/{id}", ctl.get)
	r.Get("/", ctl.list)

	mux.Mount("/pets/v1", r)
}

func getPetIDParameter(r *http.Request) (string, error) {
	id, err := urlparam.GetUUIDParameter(r, petIDParameter)
	if err != nil {
		return "", err
	}

	return id, nil
}
