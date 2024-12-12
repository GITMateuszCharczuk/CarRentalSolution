package repository

import (
	repository_interfaces "identity-api/Domain/repository_interfaces/refresh_token_repository"
	redis_config "identity-api/Infrastructure/databases/redis/config"

	"github.com/google/wire"
)

func ProvideRefreshTokenCommandRepository(redisConfig *redis_config.RedisDatabase) repository_interfaces.RefreshTokenCommandRepository {
	return NewRefreshTokenCommandRepositoryImpl(redisConfig.Client, redisConfig.Ctx)
}

func ProvideRefreshTokenQueryRepository(redisConfig *redis_config.RedisDatabase) repository_interfaces.RefreshTokenQueryRepository {
	return NewRefreshTokenQueryRepositoryImpl(redisConfig.Client, redisConfig.Ctx)
}

var WireSet = wire.NewSet(ProvideRefreshTokenCommandRepository, ProvideRefreshTokenQueryRepository)
