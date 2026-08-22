package purchase

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// PurchaseRedis represents the Redis caching repository for the purchase module.
type PurchaseRedis struct {
	client *redis.Client
}

// NewPurchaseRedis initializes a new Redis cache instance.
func NewPurchaseRedis(client *redis.Client) *PurchaseRedis {
	return &PurchaseRedis{client: client}
}

// GetVendorCatalog : Caching daftar pemasok dan katalog historis
func (r *PurchaseRedis) GetVendorCatalog(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
