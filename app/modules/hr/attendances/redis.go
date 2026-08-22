package attendances

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// AttendancesRedis represents the Redis caching repository for the attendances module.
type AttendancesRedis struct {
	client *redis.Client
}

// NewAttendancesRedis initializes a new Redis cache instance.
func NewAttendancesRedis(client *redis.Client) *AttendancesRedis {
	return &AttendancesRedis{client: client}
}

// BufferAttendance : Buffer antrean clock-in/out harian
func (r *AttendancesRedis) BufferAttendance(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
