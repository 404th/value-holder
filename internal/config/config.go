package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ProjectHost string `default:"localhost"`
	ProjectPort uint8  `default:"8050"`
	ProjectMode string `default:"development"` // ['development', 'production']

	PostgresHost                    string `default:"0.0.0.0"`
	PostgresPort                    uint8  `default:"5432"`
	PostgresDatabase                string `default:"postgres"`
	PostgresUser                    string `default:"postgres"`
	PostgresPassword                string `default:"postgres"`
	PostgresMaxConnections          int32  `default:"50"`
	PostgresSSLMode                 string `default:"disable"`
	PostgresDockerContainerName     string `default:"psql_container"`
	PostgresDockerContainerUser     string `default:"postgres"`
	PostgresDockerContainerPassword string `default:"postgres"`

	RedisHost                    string `default:"0.0.0.0"`
	RedisPort                    uint8  `default:"5432"`
	RedisDatabase                string `default:"root"`
	RedisUser                    string `default:"root"`
	RedisPassword                string `default:"root"`
	RedisMaxConnections          int32  `default:"50"`
	RedisDockerContainerName     string `default:"psql_container"`
	RedisDockerContainerUser     string `default:"root"`
	RedisDockerContainerPassword string `default:"root"`
}

func NewConfig() (cfg *Config, err error) {
	cfg = new(Config)

	if err = godotenv.Overload(EnvFileDevelopment); err != nil {
		return
	}

	cfg.ProjectHost = cast.ToString(getDotEnvValueIfExistOrReturnDefault("ProjectHost", "localhost"))
	cfg.ProjectPort = cast.ToUint8(getDotEnvValueIfExistOrReturnDefault("ProjectPort", 8050))
	cfg.ProjectMode = cast.ToString(getDotEnvValueIfExistOrReturnDefault("ProjectMode", "development"))

	cfg.PostgresHost = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresHost", "0.0.0.0"))
	cfg.PostgresPort = cast.ToUint8(getDotEnvValueIfExistOrReturnDefault("PostgresPort", 5432))
	cfg.PostgresDatabase = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresDatabase", "postgres"))
	cfg.PostgresUser = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresUser", "postgres"))
	cfg.PostgresPassword = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresPassword", "postgres"))
	cfg.PostgresMaxConnections = cast.ToInt32(getDotEnvValueIfExistOrReturnDefault("PostgresMaxConnections", 100))
	cfg.PostgresSSLMode = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresSSLMode", "disable"))
	cfg.PostgresDockerContainerName = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresDockerContainerName", "postgres_docker"))
	cfg.PostgresDockerContainerUser = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresDockerContainerUser", "postgres"))
	cfg.PostgresDockerContainerPassword = cast.ToString(getDotEnvValueIfExistOrReturnDefault("PostgresDockerContainerPassword", "super_secret_password"))

	cfg.RedisHost = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisHost", "localhost"))
	cfg.RedisPort = cast.ToUint8(getDotEnvValueIfExistOrReturnDefault("RedisPort", 6379))
	cfg.RedisDatabase = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisDatabase", "redis_database"))
	cfg.RedisUser = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisUser", "redis_user"))
	cfg.RedisPassword = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisPassword", "super_secret_redis_password"))
	cfg.RedisMaxConnections = cast.ToInt32(getDotEnvValueIfExistOrReturnDefault("RedisMaxConnections", 100))
	cfg.RedisDockerContainerName = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisDockerContainerName", "redis_docker"))
	cfg.RedisDockerContainerUser = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisDockerContainerUser", "root"))
	cfg.RedisDockerContainerPassword = cast.ToString(getDotEnvValueIfExistOrReturnDefault("RedisDockerContainerPassword", "root"))

	return
}

func getDotEnvValueIfExistOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if os.Getenv(key) != "" {
		return key
	}

	return defaultValue
}
