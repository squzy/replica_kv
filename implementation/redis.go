package implementation

import (
	"github.com/go-redis/redis"
	"github.com/squzy/replica_kv"
	"time"
)

type redis_kv struct {
	client *redis.Client
}

func (r *redis_kv) Set(key, value string, duration time.Duration) error {
	return r.client.Set(key, value, duration).Err()
}

func (r *redis_kv) Get(key string) (string, error) {
	return r.client.Get( key).Result()
}

func New(opts *redis.Options) replica_kv.ReplicaKv {
	return &redis_kv{
		client: redis.NewClient(opts),
	}
}