package inventory

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// InventoryRedis represents the Redis caching repository for the inventory module.
type InventoryRedis struct {
	client *redis.Client
}

// NewInventoryRedis initializes a new Redis cache instance.
func NewInventoryRedis(client *redis.Client) *InventoryRedis {
	return &InventoryRedis{client: client}
}

// GetTotalAvailableStock : Mengambil angka final sisa barang untuk eCommerce
func (r *InventoryRedis) GetTotalAvailableStock(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
