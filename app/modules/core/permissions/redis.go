package permissions

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// PermissionsRedis represents the Redis caching repository for the permissions module.
type PermissionsRedis struct {
	client *redis.Client
}

// NewPermissionsRedis initializes a new Redis cache instance.
func NewPermissionsRedis(client *redis.Client) *PermissionsRedis {
	return &PermissionsRedis{client: client}
}

// GetPermission : Caching hak akses spesifik rute API
func (r *PermissionsRedis) GetPermission(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
