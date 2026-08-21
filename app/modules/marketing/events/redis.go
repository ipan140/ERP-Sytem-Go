package events

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// EventsRedis represents the Redis caching repository for the events module.
type EventsRedis struct {
	client *redis.Client
}

// NewEventsRedis initializes a new Redis cache instance.
func NewEventsRedis(client *redis.Client) *EventsRedis {
	return &EventsRedis{client: client}
}

// DecrementTicketQuota : Mengurangi sisa kuota tiket secara atomic (Flash Sale)
func (r *EventsRedis) DecrementTicketQuota(ctx context.Context, key string) (int64, error) {
	return r.client.Decr(ctx, key).Result()
}

// GetEventDetails : Mengambil detail event yang sedang berjalan
func (r *EventsRedis) GetEventDetails(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
