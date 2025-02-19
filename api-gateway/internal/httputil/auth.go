package httputil

import (
	"fmt"
	"net/http"
)

type userIDContextKey = int

const (
	UserIDContextKey userIDContextKey = iota
)

func ExtractUserIDFromCtx(r *http.Request) (string, error) {
	userID, ok := r.Context().Value(UserIDContextKey).(string)
	if !ok {
		return "", fmt.Errorf("user_id not found in context or is not a string")
	}

	if userID == "" {
		return "", fmt.Errorf("user_id is empty")
	}
	return userID, nil
}
