package ecommerce

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

// EcommerceRedis represents the Redis caching repository for the ecommerce module.
type EcommerceRedis struct {
	client *redis.Client
}

// NewEcommerceRedis initializes a new Redis cache instance.
func NewEcommerceRedis(client *redis.Client) *EcommerceRedis {
	return &EcommerceRedis{client: client}
}

// GetProductCatalog : Mengambil daftar katalog produk eCommerce
func (r *EcommerceRedis) GetProductCatalog(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}

// SetProductCatalog : Menyimpan daftar katalog produk
func (r *EcommerceRedis) SetProductCatalog(ctx context.Context, key string, data interface{}, expiration time.Duration) error {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, bytesData, expiration).Err()
}

// GetShoppingCartSession : Mengambil keranjang belanja sementara pelanggan
func (r *EcommerceRedis) GetShoppingCartSession(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}

// SetShoppingCartSession : Menyimpan keranjang belanja pelanggan
func (r *EcommerceRedis) SetShoppingCartSession(ctx context.Context, key string, data interface{}, expiration time.Duration) error {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, bytesData, expiration).Err()
}
