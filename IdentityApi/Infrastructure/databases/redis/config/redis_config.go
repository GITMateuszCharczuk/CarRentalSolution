package redis_config

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Client *redis.Client
	Ctx    context.Context
}

var (
	instance *RedisConfig
	once     sync.Once
)

func NewRedisConfig(redisHost, redisPort, redisPassword string) (*RedisConfig, error) {
	var err error
	once.Do(func() {
		instance, err = initializeRedisConfig(redisHost, redisPort, redisPassword)
	})
	return instance, err
}

func initializeRedisConfig(redisHost, redisPort, redisPassword string) (*RedisConfig, error) {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
		DB:       0,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("could not connect to Redis: %v", err)
	}

	return &RedisConfig{
		Client: client,
		Ctx:    ctx,
	}, nil
}

func (r *RedisConfig) Close() error {
	return r.Client.Close()
}
