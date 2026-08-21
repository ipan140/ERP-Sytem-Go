package mass_mailing

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Mass_mailingRedis represents the Redis caching repository for the mass_mailing module.
type Mass_mailingRedis struct {
	client *redis.Client
}

// NewMass_mailingRedis initializes a new Redis cache instance.
func NewMass_mailingRedis(client *redis.Client) *Mass_mailingRedis {
	return &Mass_mailingRedis{client: client}
}

// GetEmailTemplate : Mengambil template HTML email massal
func (r *Mass_mailingRedis) GetEmailTemplate(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
