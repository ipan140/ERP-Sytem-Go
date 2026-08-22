package subscriptions

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// SubscriptionsRedis represents the Redis caching repository for the subscriptions module.
type SubscriptionsRedis struct {
	client *redis.Client
}

// NewSubscriptionsRedis initializes a new Redis cache instance.
func NewSubscriptionsRedis(client *redis.Client) *SubscriptionsRedis {
	return &SubscriptionsRedis{client: client}
}

// GetSubscriptionStatus : Caching status langganan aktif/expired
func (r *SubscriptionsRedis) GetSubscriptionStatus(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
