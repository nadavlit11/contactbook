package database

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"sync"
)

var redisClientOnce sync.Once
var redisClient RedisClient
var rConn *redis.Client

type RedisClient interface {
	GetConn() *redis.Client
}

type RedisClientImpl struct {
}

func NewRedisClient() RedisClient {
	redisClientOnce.Do(func() {
		redisClient = &RedisClientImpl{}
		rConn = redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		})
		err := rConn.Set(context.Background(), "foo", "bar", 0).Err()
		if err != nil {
			panic(err)
		}
	})
	return redisClient
}

func (p *RedisClientImpl) GetConn() *redis.Client {
	return rConn
}
