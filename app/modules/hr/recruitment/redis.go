package recruitment

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// RecruitmentRedis represents the Redis caching repository for the recruitment module.
type RecruitmentRedis struct {
	client *redis.Client
}

// NewRecruitmentRedis initializes a new Redis cache instance.
func NewRecruitmentRedis(client *redis.Client) *RecruitmentRedis {
	return &RecruitmentRedis{client: client}
}

// GetJobPostings : Mengambil daftar lowongan kerja aktif untuk website publik
func (r *RecruitmentRedis) GetJobPostings(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}
