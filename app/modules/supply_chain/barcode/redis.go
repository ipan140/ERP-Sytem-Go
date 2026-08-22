package barcode

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// BarcodeRedis represents the Redis caching repository for the barcode module.
type BarcodeRedis struct {
	client *redis.Client
}

// NewBarcodeRedis initializes a new Redis cache instance.
func NewBarcodeRedis(client *redis.Client) *BarcodeRedis {
	return &BarcodeRedis{client: client}
}

// GetBarcodeMapping : Caching pencocokan nomor barcode ke data Produk
func (r *BarcodeRedis) GetBarcodeMapping(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
