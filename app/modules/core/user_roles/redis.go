package user_roles

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// UserRolesRedis represents the Redis caching repository for the user_roles module.
type UserRolesRedis struct {
	client *redis.Client
}

// NewUserRolesRedis initializes a new Redis cache instance.
func NewUserRolesRedis(client *redis.Client) *UserRolesRedis {
	return &UserRolesRedis{client: client}
}

// GetUserRole : Caching peran pengguna (misal: Admin, Staf, Manajer)
func (r *UserRolesRedis) GetUserRole(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
