package blog

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// BlogRedis represents the Redis caching repository for the blog module.
type BlogRedis struct {
	client *redis.Client
}

// NewBlogRedis initializes a new Redis cache instance.
func NewBlogRedis(client *redis.Client) *BlogRedis {
	return &BlogRedis{client: client}
}

// GetBlogPosts : Mengambil daftar artikel blog statis
func (r *BlogRedis) GetBlogPosts(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
