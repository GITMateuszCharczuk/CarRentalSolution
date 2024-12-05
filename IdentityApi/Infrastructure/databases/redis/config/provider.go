package redis_config

import (
	"identity-api/Infrastructure/config"

	"github.com/google/wire"
)

func NewRedisConfigProvider(cfg *config.Config) (*RedisDatabase, error) {
	return NewRedisConfig(cfg.RedisHost, cfg.RedisPort, cfg.RedisPassword)
}

var WireSet = wire.NewSet(NewRedisConfigProvider)
