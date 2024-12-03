package repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RefreshTokenQueryRepositoryImpl struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRefreshTokenQueryRepositoryImpl(client *redis.Client, ctx context.Context) *RefreshTokenQueryRepositoryImpl {
	return &RefreshTokenQueryRepositoryImpl{
		Client: client,
		Ctx:    ctx,
	}
}

func (r *RefreshTokenQueryRepositoryImpl) GetRefreshToken(refreshToken string) (string, error) {
	ctx := r.Ctx
	userID, err := r.Client.Get(ctx, refreshToken).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("refresh token not found: %s", refreshToken)
		}
		return "", fmt.Errorf("could not get user ID from refresh token: %v", err)
	}
	return userID, nil
}
