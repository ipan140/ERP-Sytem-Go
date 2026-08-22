package timesheets

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// TimesheetsRedis represents the Redis caching repository for the timesheets module.
type TimesheetsRedis struct {
	client *redis.Client
}

// NewTimesheetsRedis initializes a new Redis cache instance.
func NewTimesheetsRedis(client *redis.Client) *TimesheetsRedis {
	return &TimesheetsRedis{client: client}
}

// GetProjectHours : Caching total jam kerja proyek
func (r *TimesheetsRedis) GetProjectHours(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
