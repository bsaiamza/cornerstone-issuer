package config

import (
	"os"
)

type Config struct {
}

func LoadConfig() *Config {
	return &Config{}
}

func getEnvVarByName(name string) string {
	return os.Getenv(name)
}

func (c *Config) GetAcapyURL() string {
	return getEnvVarByName("ACAPY_URL")
}

func (c *Config) GetDHAAPI() string {
	return getEnvVarByName("DHAAPI")
}

func (c *Config) GetSchemaID() string {
	return getEnvVarByName("SCHEMA_ID")
}

func (c *Config) GetSchemaName() string {
	return getEnvVarByName("SCHEMA_NAME")
}

func (c *Config) GetSchemaVersion() string {
	return getEnvVarByName("SCHEMA_VERSION")
}

func (c *Config) GetCredDefID() string {
	return getEnvVarByName("CRED_DEF_ID")
}

func (c *Config) GetPublicDID() string {
	return getEnvVarByName("PUBLIC_DID")
}

func (c *Config) GetServerAddress() string {
	return getEnvVarByName("SERVER_ADDRESS")
}
