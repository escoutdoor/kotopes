package token

import (
	"fmt"
	"time"

	"github.com/escoutdoor/kotopes/auth/internal/config"
	"github.com/escoutdoor/kotopes/auth/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

type Provider interface {
	Verify(token string) (*model.UserClaims, error)
	GenerateToken(id, role string) (string, error)
}

type tokenProvider struct {
	cfg config.TokenConfig
}

func NewTokenProvider(cfg config.TokenConfig) Provider {
	return &tokenProvider{
		cfg: cfg,
	}
}

func (p *tokenProvider) GenerateToken(id, role string) (string, error) {
	claims := &model.UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(p.cfg.TTL())),
		},
		ID:   id,
		Role: role,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(p.cfg.SecretKey())
	if err != nil {
		return "", fmt.Errorf("sign jwt token: %s", err)
	}
	return token, nil
}

func (p *tokenProvider) Verify(token string) (*model.UserClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected token signing method")
		}

		return p.cfg.SecretKey(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := jwtToken.Claims.(*model.UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
