package repository

import (
	"identity-api/Domain/repository_interfaces"
	redis_config "identity-api/Infrastructure/databases/redis/config"

	"github.com/google/wire"
)

func ProvideRefreshTokenRepository(redisConfig *redis_config.RedisConfig) repository_interfaces.RefreshTokenRepository {
	return NewRefreshTokenRepository(redisConfig.Client, redisConfig.Ctx)
}

var ProviderSet = wire.NewSet(NewRefreshTokenRepository)
