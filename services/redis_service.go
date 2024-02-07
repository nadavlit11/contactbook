package services

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

var redisServiceOnce sync.Once
var redisService RedisService

type RedisService interface {
	GetConn() *redis.Client
}

type RedisServiceImpl struct {
}

func NewRedisService() RedisService {
	redisServiceOnce.Do(func() {
		redisService = &RedisServiceImpl{}
	})
	return redisService
}

func (service *RedisServiceImpl) GetConn() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
