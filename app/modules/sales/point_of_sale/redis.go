package point_of_sale

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Point_of_saleRedis represents the Redis caching repository for the point_of_sale module.
type Point_of_saleRedis struct {
	client *redis.Client
}

// NewPoint_of_saleRedis initializes a new Redis cache instance.
func NewPoint_of_saleRedis(client *redis.Client) *Point_of_saleRedis {
	return &Point_of_saleRedis{client: client}
}

// GetPOSCatalog : Mengambil katalog produk instan untuk kasir offline
func (r *Point_of_saleRedis) GetPOSCatalog(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
