package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	grpc_client "github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc"
	"github.com/escoutdoor/kotopes/api-gateway/internal/httputil"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const (
	authorizationHeader       = "Authorization"
	authorizationHeaderPrefix = "Bearer "
)

func AuthMiddleware(
	authServiceClient grpc_client.AuthClient,
	accessServiceClient grpc_client.AccessClient,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get(authorizationHeader)
			if authHeader == "" {
				render.Render(w, r, httputil.ErrUnauthorized(fmt.Errorf("authorization header not provided")))
				return
			}
			if !strings.HasPrefix(authHeader, authorizationHeaderPrefix) {
				render.Render(w, r, httputil.ErrUnauthorized(fmt.Errorf("invalid authorization header format")))
				return
			}
			token := authHeader[len(authorizationHeaderPrefix):]

			tokenData, err := authServiceClient.Validate(r.Context(), token)
			if err != nil {
				httputil.HandleGrpcError(w, r, err)
				return
			}

			rctx := chi.RouteContext(r.Context())
			isAllowed, err := accessServiceClient.CheckIsAllowed(r.Context(), &model.AccessCheck{
				Endpoint: rctx.RoutePatterns[0],
				Method:   r.Method,
				UserID:   tokenData.UserID,
				Role:     tokenData.Role,
			})
			if err != nil {
				httputil.HandleGrpcError(w, r, err)
				return
			}
			if !isAllowed {
				render.Render(w, r, httputil.ErrForbidden(fmt.Errorf("you do not have access to this resource")))
				return
			}

			ctx := context.WithValue(r.Context(), httputil.UserIDContextKey, tokenData.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
