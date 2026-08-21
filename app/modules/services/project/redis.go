package project

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// ProjectRedis represents the Redis caching repository for the project module.
type ProjectRedis struct {
	client *redis.Client
}

// NewProjectRedis initializes a new Redis cache instance.
func NewProjectRedis(client *redis.Client) *ProjectRedis {
	return &ProjectRedis{client: client}
}

// GetGanttChart : Mengambil ringkasan milestone untuk grafik Gantt timeline
func (r *ProjectRedis) GetGanttChart(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
