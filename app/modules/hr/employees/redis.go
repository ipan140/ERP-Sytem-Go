package employees

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// EmployeesRedis represents the Redis caching repository for the employees module.
type EmployeesRedis struct {
	client *redis.Client
}

// NewEmployeesRedis initializes a new Redis cache instance.
func NewEmployeesRedis(client *redis.Client) *EmployeesRedis {
	return &EmployeesRedis{client: client}
}

// GetPhonebook : Mengambil direktori kontak karyawan
func (r *EmployeesRedis) GetPhonebook(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}

// GetHolidays : Mengambil kalender hari libur nasional
func (r *EmployeesRedis) GetHolidays(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
