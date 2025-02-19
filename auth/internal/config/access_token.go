package config

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

var (
	accessTokenTTLEnvName       = "ACCESS_TOKEN_TTL"
	accessTokenSecretKeyEnvName = "ACCESS_TOKEN_SECRET_KEY"
)

type accessTokenConfig struct {
	accessTokenTTL time.Duration
	secretKey      []byte
}

func NewAccessTokenConfig() (TokenConfig, error) {
	accessTokenTTLStr := os.Getenv(accessTokenTTLEnvName)
	if accessTokenTTLStr == "" {
		return nil, fmt.Errorf("access token ttl is empty")
	}

	accessTokenTTL, err := time.ParseDuration(accessTokenTTLStr)
	if err != nil {
		return nil, fmt.Errorf("parse access token ttl error: %s", err)
	}

	secretKey := os.Getenv(accessTokenSecretKeyEnvName)
	if secretKey == "" {
		return nil, fmt.Errorf("jwt secret key is empty")
	}
	decodedSecretKey, err := base64.StdEncoding.DecodeString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("decode access token secret key error: %s", err)
	}

	return &accessTokenConfig{
		accessTokenTTL: accessTokenTTL,
		secretKey:      decodedSecretKey,
	}, nil
}

func (c *accessTokenConfig) SecretKey() []byte {
	return c.secretKey
}

func (c *accessTokenConfig) TTL() time.Duration {
	return c.accessTokenTTL
}
