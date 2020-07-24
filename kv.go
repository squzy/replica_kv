package replica_kv

import "time"

type ReplicaKv interface {
	Set(key, value string, duration time.Duration) (error)
	Get(key string) (string, error)
}