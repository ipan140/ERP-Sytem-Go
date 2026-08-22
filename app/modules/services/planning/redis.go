package planning

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// PlanningRedis represents the Redis caching repository for the planning module.
type PlanningRedis struct {
	client *redis.Client
}

// NewPlanningRedis initializes a new Redis cache instance.
func NewPlanningRedis(client *redis.Client) *PlanningRedis {
	return &PlanningRedis{client: client}
}

// GetAvailableSlots : Caching slot kalender
func (r *PlanningRedis) GetAvailableSlots(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
