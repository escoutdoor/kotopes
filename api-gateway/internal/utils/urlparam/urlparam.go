package urlparam

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetUUIDParameter(r *http.Request, parameter string) (string, error) {
	id := chi.URLParam(r, parameter)
	if id == "" {
		return "", fmt.Errorf("id parameter should not be empty")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return "", fmt.Errorf("id parameter must be a valid uuid")
	}

	return id, nil
}
