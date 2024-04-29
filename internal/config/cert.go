package config

import (
	"errors"
	"os"
)

const (
	certFileEnvName    = "CERT_FILE"
	certKeyFileEnvName = "CERT_KEY_FILE"
)

var (
	ErrCertFileNotFound    = errors.New("cert file not found")
	ErrCertKeyFileNotFound = errors.New("cert key file not found")
)

type CERTConfig interface {
	CertFile() string
	CertKeyFile() string
}

type certConfig struct {
	certFile    string
	certKeyFile string
}

func NewCERTConfig() (CERTConfig, error) {
	certFile := os.Getenv(certFileEnvName)
	if len(certFile) == 0 {
		return nil, ErrCertFileNotFound
	}

	certKeyFile := os.Getenv(certKeyFileEnvName)
	if len(certKeyFile) == 0 {
		return nil, ErrCertKeyFileNotFound
	}

	return &certConfig{
		certFile:    certFile,
		certKeyFile: certKeyFile,
	}, nil
}

func (cfg *certConfig) CertFile() string {
	return cfg.certFile
}

func (cfg *certConfig) CertKeyFile() string {
	return cfg.certKeyFile
}
