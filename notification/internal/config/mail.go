package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	mailHostEnvName     = "MAIL_HOST"
	mailPortEnvName     = "MAIL_PORT"
	mailUsernameEnvName = "MAIL_USERNAME"
	mailPasswordEnvName = "MAIL_PASSWORD"
)

type mailConfig struct {
	host string
	port int

	username string
	password string
}

func NewMailConfig() (MailConfig, error) {
	host := os.Getenv(mailHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("mail host is not defined or empty")
	}

	portStr := os.Getenv(mailPortEnvName)
	if portStr == "" {
		return nil, fmt.Errorf("mail port is not defined or empty")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid mail port: %s", err)
	}

	username := os.Getenv(mailUsernameEnvName)
	if username == "" {
		return nil, fmt.Errorf("mail username is not defined or empty")
	}

	password := os.Getenv(mailPasswordEnvName)
	if password == "" {
		return nil, fmt.Errorf("mail password is not defined or empty")
	}

	return &mailConfig{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}, nil
}

func (c *mailConfig) Host() string {
	return c.host
}

func (c *mailConfig) Port() int {
	return c.port
}

func (c *mailConfig) From() string {
	return c.username
}

func (c *mailConfig) Password() string {
	return c.password
}
