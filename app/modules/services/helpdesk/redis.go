package helpdesk

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// HelpdeskRedis represents the Redis caching repository for the helpdesk module.
type HelpdeskRedis struct {
	client *redis.Client
}

// NewHelpdeskRedis initializes a new Redis cache instance.
func NewHelpdeskRedis(client *redis.Client) *HelpdeskRedis {
	return &HelpdeskRedis{client: client}
}

// GetSLARules : Mengambil aturan SLA maksimum respons
func (r *HelpdeskRedis) GetSLARules(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}

// GetCannedResponses : Mengambil template jawaban FAQ otomatis
func (r *HelpdeskRedis) GetCannedResponses(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
