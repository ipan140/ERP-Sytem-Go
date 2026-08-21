package elearning

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// ElearningRedis represents the Redis caching repository for the elearning module.
type ElearningRedis struct {
	client *redis.Client
}

// NewElearningRedis initializes a new Redis cache instance.
func NewElearningRedis(client *redis.Client) *ElearningRedis {
	return &ElearningRedis{client: client}
}

// GetCourseCatalog : Mengambil daftar silabus kursus statis
func (r *ElearningRedis) GetCourseCatalog(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
