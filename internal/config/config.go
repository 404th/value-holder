package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ValueHolderProjectHost string
	ValueHolderProjectPort uint32
	ValueHolderProjectMode string

	PostgresHost                    string
	PostgresPort                    uint32
	PostgresDatabase                string
	PostgresUser                    string
	PostgresPassword                string
	PostgresMaxConnections          uint32
	PostgresSSLMode                 string
	PostgresDockerContainerName     string
	PostgresDockerContainerUser     string
	PostgresDockerContainerPassword string

	RedisHost                    string
	RedisPort                    uint8
	RedisDatabase                string
	RedisUser                    string
	RedisPassword                string
	RedisMaxConnections          uint32
	RedisDockerContainerName     string
	RedisDockerContainerUser     string
	RedisDockerContainerPassword string
}

func NewConfig() (cfg *Config, err error) {
	cfg = new(Config)

	if err = godotenv.Overload(EnvFileDevelopment); err != nil {
		return
	}

	cfg.ValueHolderProjectHost = cast.ToString(getDotEnvValueIfExistOrReturnDefault("ValueHolderProjectHost", "localhost"))
	cfg.ValueHolderProjectPort = cast.ToUint32(getDotEnvValueIfExistOrReturnDefault("ValueHolderProjectPort", 8050))
	cfg.ValueHolderProjectMode = cast.ToString(getDotEnvValueIfExistOrReturnDefault("ValueHolderProjectMode", "development"))

	cfg.PostgresHost = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresHost", "0.0.0.0"))
	cfg.PostgresPort = cast.ToUint32(getDotEnvValueIfExistOrReturnDefault("PostgresPort", 5432))
	cfg.PostgresDatabase = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresDatabase", "postgres"))
	cfg.PostgresUser = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresUser", "postgres"))
	cfg.PostgresPassword = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresPassword", "postgres"))
	cfg.PostgresMaxConnections = cast.ToUint32(getDotEnvValueIfExistOrReturnDefault("PostgresMaxConnections", 100))
	cfg.PostgresSSLMode = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresSSLMode", "disable"))
	cfg.PostgresDockerContainerName = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresDockerContainerName", "value_holder_postgresql"))
	cfg.PostgresDockerContainerUser = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresDockerContainerUser", "postgres"))
	cfg.PostgresDockerContainerPassword = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresDockerContainerPassword", "super_secret_password"))

	cfg.RedisHost = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisHost", "localhost"))
	cfg.RedisPort = cast.ToUint8(getDotEnvValueIfExistOrReturnDefault("RedisPort", 6379))
	cfg.RedisDatabase = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisDatabase", "redis_database"))
	cfg.RedisUser = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisUser", "redis_user"))
	cfg.RedisPassword = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisPassword", "super_secret_redis_password"))
	cfg.RedisMaxConnections = cast.ToUint32(getDotEnvValueIfExistOrReturnDefault("RedisMaxConnections", 100))
	cfg.RedisDockerContainerName = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisDockerContainerName", "redis_docker"))
	cfg.RedisDockerContainerUser = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisDockerContainerUser", "root"))
	cfg.RedisDockerContainerPassword = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisDockerContainerPassword", "root"))

	return
}

func getDotEnvValueIfExistOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}

	return defaultValue
}
