package invoicing

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// InvoicingRedis represents the Redis caching repository for the invoicing module.
type InvoicingRedis struct {
	client *redis.Client
}

// NewInvoicingRedis initializes a new Redis cache instance.
func NewInvoicingRedis(client *redis.Client) *InvoicingRedis {
	return &InvoicingRedis{client: client}
}

// GetDashboardStats : Caching total piutang/hutang bulan ini
func (r *InvoicingRedis) GetDashboardStats(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
