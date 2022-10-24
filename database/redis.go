package database

import "github.com/go-redis/redis/v8"

func NewRedis(opt *redis.Options) *redis.Client {
	return redis.NewClient(opt)
}
