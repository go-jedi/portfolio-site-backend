package config

import (
	"errors"
	"os"
	"strconv"
)

const (
	isTLSEnvName       = "IS_TLS"
	certFileEnvName    = "CERT_FILE"
	certKeyFileEnvName = "CERT_KEY_FILE"
)

var (
	ErrIsTLSNotFound       = errors.New("is tls not found")
	ErrCertFileNotFound    = errors.New("cert file not found")
	ErrCertKeyFileNotFound = errors.New("cert key file not found")
)

type CERTConfig interface {
	IsTLS() bool
	CertFile() string
	CertKeyFile() string
}

type certConfig struct {
	isTLS       bool
	certFile    string
	certKeyFile string
}

func NewCERTConfig() (CERTConfig, error) {
	isTLS := os.Getenv(isTLSEnvName)
	if len(isTLS) == 0 {
		return nil, ErrIsTLSNotFound
	}

	isTLSBool, err := strconv.ParseBool(isTLS)
	if err != nil {
		return nil, err
	}

	certFile := os.Getenv(certFileEnvName)
	if len(certFile) == 0 {
		return nil, ErrCertFileNotFound
	}

	certKeyFile := os.Getenv(certKeyFileEnvName)
	if len(certKeyFile) == 0 {
		return nil, ErrCertKeyFileNotFound
	}

	return &certConfig{
		isTLS:       isTLSBool,
		certFile:    certFile,
		certKeyFile: certKeyFile,
	}, nil
}

func (cfg *certConfig) IsTLS() bool {
	return cfg.isTLS
}

func (cfg *certConfig) CertFile() string {
	return cfg.certFile
}

func (cfg *certConfig) CertKeyFile() string {
	return cfg.certKeyFile
}
