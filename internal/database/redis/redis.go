package redis

import (
	"fmt"

	"github.com/404th/value-holder/internal/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func NewRedis(cfg *config.Config, logger zap.Logger) (rd *redis.Client, err error) {
	rd = new(redis.Client)

	opt, err := redis.ParseURL(fmt.Sprintf("redis://%s:%s@%s:%d/%s",
		cfg.RedisUser,
		cfg.RedisPassword,
		cfg.RedisHost,
		cfg.RedisPort,
		cfg.RedisDatabase,
	))
	if err != nil {

	}

	client := redis.NewClient(opt)

	return
}
