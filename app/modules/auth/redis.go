package auth

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

// AuthRedis represents the Redis caching repository for the auth module.
type AuthRedis struct {
	client *redis.Client
}

// NewAuthRedis initializes a new Redis cache instance.
func NewAuthRedis(client *redis.Client) *AuthRedis {
	return &AuthRedis{client: client}
}

// SetJWTBlacklist : Menyimpan token JWT yang sudah di-logout
func (r *AuthRedis) SetJWTBlacklist(ctx context.Context, key string, data interface{}, expiration time.Duration) error {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, bytesData, expiration).Err()
}

// GetRBAC : Mengambil data Roles & Permissions dari cache
func (r *AuthRedis) GetRBAC(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}

// SetRBAC : Menyimpan data Roles & Permissions ke cache
func (r *AuthRedis) SetRBAC(ctx context.Context, key string, data interface{}, expiration time.Duration) error {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, bytesData, expiration).Err()
}
