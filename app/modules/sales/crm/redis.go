package crm

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// CrmRedis represents the Redis caching repository for the crm module.
type CrmRedis struct {
	client *redis.Client
}

// NewCrmRedis initializes a new Redis cache instance.
func NewCrmRedis(client *redis.Client) *CrmRedis {
	return &CrmRedis{client: client}
}

// GetKanbanBoard : Caching papan Kanban Daftar Prospek
func (r *CrmRedis) GetKanbanBoard(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
