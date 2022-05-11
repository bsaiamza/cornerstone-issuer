package config

import (
	"cornerstone_issuer/pkg/log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config stores all env vars.
type Config struct {
}

// GetConfig returns the config.
func GetConfig() *Config {
	config := &Config{}

	envFilePath, _ := filepath.Abs(".env")

	if err := godotenv.Load(envFilePath); err != nil {
		log.ServerWarning.Print("No .env file found, using env vars from os.")
	}

	return config
}

// getEnv loads env vars from .env file or os.
func getEnv(key string) string {
	return os.Getenv(key)
}

// GetAcapyURL returns the Acapy URL.
func (c *Config) GetAcapyURL() string {
	return getEnv("ACAPY_URL")
}

// GetAcapyAdminAPIKey returns the Acapy Admin API key.
func (c *Config) GetAcapyAdminAPIKey() string {
	return getEnv("ACAPY_ADMIN_API_KEY")
}

// GetClientURL returns the client URL.
func (c *Config) GetClientURL() string {
	return getEnv("CLIENT_URL")
}

// GetServerHost returns the server host.
func (c *Config) GetServerHost() string {
	return getEnv("SERVER_HOST")
}

// GetServerPort returns the server port.
func (c *Config) GetServerPort() string {
	return getEnv("SERVER_PORT")
}