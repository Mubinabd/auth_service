package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	Environment      string
	LogLevel         string

	UserServiceHost string
	UserServicePort string

	PostServiceHost string
	PostServicePort string


}

func Load() Config {
	c := Config{}

	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "ahrorbek"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "ahrorbek"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "userdb"))
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	c.UserServicePort = cast.ToString(getOrReturnDefault("USER_SERVICE_PORT", ":8001"))

	c.PostServiceHost = cast.ToString(getOrReturnDefault("PRODUCT_SERVICE_HOST", "localhost"))
	c.PostServicePort = cast.ToString(getOrReturnDefault("PRODUCT_SERVICE_PORT", ":8002"))

	
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
