package time_off

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// TimeOffRedis represents the Redis caching repository for the time_off module.
type TimeOffRedis struct {
	client *redis.Client
}

// NewTimeOffRedis initializes a new Redis cache instance.
func NewTimeOffRedis(client *redis.Client) *TimeOffRedis {
	return &TimeOffRedis{client: client}
}

// GetLeaveBalance : Caching sisa kuota cuti karyawan
func (r *TimeOffRedis) GetLeaveBalance(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
