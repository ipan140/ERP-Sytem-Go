package base

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// BaseRedis represents the Redis caching repository for the base module.
type BaseRedis struct {
	client *redis.Client
}

// NewBaseRedis initializes a new Redis cache instance.
func NewBaseRedis(client *redis.Client) *BaseRedis {
	return &BaseRedis{client: client}
}

// GetCompanySettings : Mengambil pengaturan global perusahaan
func (r *BaseRedis) GetCompanySettings(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}

// GetCurrencyRates : Mengambil nilai tukar mata uang harian
func (r *BaseRedis) GetCurrencyRates(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}

// GetI18n : Mengambil kamus bahasa terjemahan
func (r *BaseRedis) GetI18n(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
