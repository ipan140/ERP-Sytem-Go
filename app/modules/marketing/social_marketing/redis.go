package social_marketing

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// SocialMarketingRedis represents the Redis caching repository for the social_marketing module.
type SocialMarketingRedis struct {
	client *redis.Client
}

// NewSocialMarketingRedis initializes a new Redis cache instance.
func NewSocialMarketingRedis(client *redis.Client) *SocialMarketingRedis {
	return &SocialMarketingRedis{client: client}
}

// GetPostTemplates : Caching template pesan promosi
func (r *SocialMarketingRedis) GetPostTemplates(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
