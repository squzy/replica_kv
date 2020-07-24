# MutileReplica KV implentations

```go
type ReplicaKv interface {
	Set(key, value string, duration time.Duration) (error)
	Get(key string) (string, error)
}
```