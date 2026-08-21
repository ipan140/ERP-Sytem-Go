package sales_core

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Sales_coreRedis represents the Redis caching repository for the sales_core module.
type Sales_coreRedis struct {
	client *redis.Client
}

// NewSales_coreRedis initializes a new Redis cache instance.
func NewSales_coreRedis(client *redis.Client) *Sales_coreRedis {
	return &Sales_coreRedis{client: client}
}

// GetPricelists : Mengambil aturan harga dan diskon khusus B2B
func (r *Sales_coreRedis) GetPricelists(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}

// GetSubscriptions : Mengambil daftar paket berlangganan
func (r *Sales_coreRedis) GetSubscriptions(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
