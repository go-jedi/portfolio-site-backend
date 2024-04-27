package config

import (
	"errors"
	"os"
	"strconv"
)

const (
	corsOriginEnvName     = "CORS_ORIGIN"
	corsMethodEnvName     = "CORS_METHOD"
	corsHeaderEnvName     = "CORS_HEADER"
	corsCredentialEnvName = "CORS_CREDENTIAL" //nolint:gosec
	corsMaxAgeEnvName     = "CORS_MAX_AGE"
)

var (
	ErrCorsOriginNotFound     = errors.New("cors origin not found")
	ErrCorsMethodNotFound     = errors.New("cors method not found")
	ErrCorsHeaderNotFound     = errors.New("cors header not found")
	ErrCorsCredentialNotFound = errors.New("cors credential not found")
	ErrCorsMaxAgeNotFound     = errors.New("cors max age not found")
)

type CORSConfig interface {
	ORIGIN() string
	METHOD() string
	HEADER() string
	CREDENTIAL() bool
	MAXAGE() int
}

type corsConfig struct {
	origin     string
	method     string
	header     string
	credential bool
	maxAge     int
}

func NewCORSConfig() (CORSConfig, error) {
	origin := os.Getenv(corsOriginEnvName)
	if len(origin) == 0 {
		return nil, ErrCorsOriginNotFound
	}

	method := os.Getenv(corsMethodEnvName)
	if len(origin) == 0 {
		return nil, ErrCorsMethodNotFound
	}

	header := os.Getenv(corsHeaderEnvName)
	if len(origin) == 0 {
		return nil, ErrCorsHeaderNotFound
	}

	credential := os.Getenv(corsCredentialEnvName)
	if len(credential) == 0 {
		return nil, ErrCorsCredentialNotFound
	}

	credentialBool, err := strconv.ParseBool(credential)
	if err != nil {
		return nil, err
	}

	maxAge := os.Getenv(corsMaxAgeEnvName)
	if len(maxAge) == 0 {
		return nil, ErrCorsMaxAgeNotFound
	}

	maxAgeInt, err := strconv.Atoi(maxAge)
	if err != nil {
		return nil, err
	}

	return &corsConfig{
		origin:     origin,
		method:     method,
		header:     header,
		credential: credentialBool,
		maxAge:     maxAgeInt,
	}, nil
}

func (cfg *corsConfig) ORIGIN() string {
	return cfg.origin
}

func (cfg *corsConfig) METHOD() string {
	return cfg.method
}

func (cfg *corsConfig) HEADER() string {
	return cfg.header
}

func (cfg *corsConfig) CREDENTIAL() bool {
	return cfg.credential
}

func (cfg *corsConfig) MAXAGE() int {
	return cfg.maxAge
}
