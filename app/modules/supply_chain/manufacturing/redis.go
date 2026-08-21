package manufacturing

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// ManufacturingRedis represents the Redis caching repository for the manufacturing module.
type ManufacturingRedis struct {
	client *redis.Client
}

// NewManufacturingRedis initializes a new Redis cache instance.
func NewManufacturingRedis(client *redis.Client) *ManufacturingRedis {
	return &ManufacturingRedis{client: client}
}

// GetBOM : Mengambil resep bahan baku perakitan pabrik (Bill of Materials)
func (r *ManufacturingRedis) GetBOM(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
