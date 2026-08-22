package appointments

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// AppointmentsRedis represents the Redis caching repository for the appointments module.
type AppointmentsRedis struct {
	client *redis.Client
}

// NewAppointmentsRedis initializes a new Redis cache instance.
func NewAppointmentsRedis(client *redis.Client) *AppointmentsRedis {
	return &AppointmentsRedis{client: client}
}

// GetAppointmentSlots : Caching jadwal teknisi
func (r *AppointmentsRedis) GetAppointmentSlots(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
