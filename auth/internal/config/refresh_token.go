package config

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

var (
	refreshTokenTTLEnvName       = "REFRESH_TOKEN_TTL"
	refreshTokenSecretKeyEnvName = "REFRESH_TOKEN_SECRET_KEY"
)

type refreshTokenConfig struct {
	refreshTokenTTL time.Duration
	secretKey       []byte
}

func NewRefreshTokenConfig() (TokenConfig, error) {
	refreshTokenTTLStr := os.Getenv(refreshTokenTTLEnvName)
	if refreshTokenTTLStr == "" {
		return nil, fmt.Errorf("refresh token ttl is empty")
	}

	refreshTokenTTL, err := time.ParseDuration(refreshTokenTTLStr)
	if err != nil {
		return nil, fmt.Errorf("parse refresh token ttl error: %s", err)
	}

	secretKey := os.Getenv(refreshTokenSecretKeyEnvName)
	if secretKey == "" {
		return nil, fmt.Errorf("jwt secret key is empty")
	}
	decodedSecretKey, err := base64.StdEncoding.DecodeString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("decode refresh token secret key error: %s", err)
	}

	return &refreshTokenConfig{
		refreshTokenTTL: refreshTokenTTL,
		secretKey:       decodedSecretKey,
	}, nil
}

func (c *refreshTokenConfig) SecretKey() []byte {
	return c.secretKey
}

func (c *refreshTokenConfig) TTL() time.Duration {
	return c.refreshTokenTTL
}
