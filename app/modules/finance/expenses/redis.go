package expenses

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// ExpensesRedis represents the Redis caching repository for the expenses module.
type ExpensesRedis struct {
	client *redis.Client
}

// NewExpensesRedis initializes a new Redis cache instance.
func NewExpensesRedis(client *redis.Client) *ExpensesRedis {
	return &ExpensesRedis{client: client}
}

// GetBudgetLimit : Caching sisa plafon budget divisi
func (r *ExpensesRedis) GetBudgetLimit(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
