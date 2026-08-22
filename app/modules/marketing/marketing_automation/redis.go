package marketing_automation

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// MarketingAutomationRedis represents the Redis caching repository for the marketing_automation module.
type MarketingAutomationRedis struct {
	client *redis.Client
}

// NewMarketingAutomationRedis initializes a new Redis cache instance.
func NewMarketingAutomationRedis(client *redis.Client) *MarketingAutomationRedis {
	return &MarketingAutomationRedis{client: client}
}

// GetCampaignSegments : Caching segmen kriteria audiens
func (r *MarketingAutomationRedis) GetCampaignSegments(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
