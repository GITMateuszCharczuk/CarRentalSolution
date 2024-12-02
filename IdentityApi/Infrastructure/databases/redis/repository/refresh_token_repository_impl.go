package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RefreshTokenRepositoryImp struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRefreshTokenRepository(client *redis.Client, ctx context.Context) *RefreshTokenRepositoryImp {
	return &RefreshTokenRepositoryImp{
		Client: client,
		Ctx:    ctx,
	}
}

func (r *RefreshTokenRepositoryImp) SaveRefreshToken(userID string, refreshToken string, ttl int) error {
	ctx := r.Ctx
	expiration := time.Duration(ttl) * 24 * time.Hour
	err := r.Client.Set(ctx, refreshToken, userID, expiration).Err()
	if err != nil {
		return fmt.Errorf("could not save refresh token: %v", err)
	}
	return nil
}

func (r *RefreshTokenRepositoryImp) GetRefreshToken(refreshToken string) (string, error) {
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

func (r *RefreshTokenRepositoryImp) RevokeRefreshToken(refreshToken string) error {
	ctx := r.Ctx
	err := r.Client.Del(ctx, refreshToken).Err()
	if err != nil {
		return fmt.Errorf("could not revoke refresh token: %v", err)
	}
	return nil
}
