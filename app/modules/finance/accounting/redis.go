package accounting

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

// AccountingRedis represents the Redis caching repository for the accounting module.
type AccountingRedis struct {
	client *redis.Client
}

// NewAccountingRedis initializes a new Redis cache instance.
func NewAccountingRedis(client *redis.Client) *AccountingRedis {
	return &AccountingRedis{client: client}
}

// SetWebhookIdempotencyLock : Mengunci Order ID dari Midtrans untuk mencegah lunas ganda
func (r *AccountingRedis) SetWebhookIdempotencyLock(ctx context.Context, key string, data interface{}, expiration time.Duration) error {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, bytesData, expiration).Err()
}

// GetMonthlyFinancialReport : Mengambil cache laporan P&L atau Balance Sheet bulanan
func (r *AccountingRedis) GetMonthlyFinancialReport(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
