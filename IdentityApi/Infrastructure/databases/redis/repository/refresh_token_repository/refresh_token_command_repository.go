package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RefreshTokenCommandRepositoryImpl struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRefreshTokenCommandRepositoryImpl(client *redis.Client, ctx context.Context) *RefreshTokenCommandRepositoryImpl {
	return &RefreshTokenCommandRepositoryImpl{
		Client: client,
		Ctx:    ctx,
	}
}

func (r *RefreshTokenCommandRepositoryImpl) SaveRefreshToken(userID string, refreshToken string, ttl int) error {
	ctx := r.Ctx
	expiration := time.Duration(ttl) * 24 * time.Hour
	err := r.Client.Set(ctx, refreshToken, userID, expiration).Err()
	if err != nil {
		return fmt.Errorf("could not save refresh token: %v", err)
	}
	return nil
}

func (r *RefreshTokenCommandRepositoryImpl) RevokeRefreshToken(refreshToken string) error {
	ctx := r.Ctx
	err := r.Client.Del(ctx, refreshToken).Err()
	if err != nil {
		return fmt.Errorf("could not revoke refresh token: %v", err)
	}
	return nil
}
