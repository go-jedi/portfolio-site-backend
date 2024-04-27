package config

import (
	"errors"
	"os"
)

const (
	fileServerDirEnvName    = "FILE_SERVER_DIR"
	fileServerPrefixEnvName = "FILE_SERVER_PREFIX"
)

var (
	ErrFileServerDirNotFound    = errors.New("file server dir not found")
	ErrFileServerPrefixNotFound = errors.New("file server prefix not found")
)

type FileServerConfig interface {
	FileServerDir() string
	FileServerPrefix() string
}

type fileServerConfig struct {
	fileServerDir    string
	fileServerPrefix string
}

func NewFileServerConfig() (FileServerConfig, error) {
	fileServerDir := os.Getenv(fileServerDirEnvName)
	if len(fileServerDir) == 0 {
		return nil, ErrFileServerDirNotFound
	}

	fileServerPrefix := os.Getenv(fileServerPrefixEnvName)
	if len(fileServerPrefix) == 0 {
		return nil, ErrFileServerPrefixNotFound
	}

	return &fileServerConfig{
		fileServerDir:    fileServerDir,
		fileServerPrefix: fileServerPrefix,
	}, nil
}

func (cfg *fileServerConfig) FileServerDir() string {
	return cfg.fileServerDir
}

func (cfg *fileServerConfig) FileServerPrefix() string {
	return cfg.fileServerPrefix
}
