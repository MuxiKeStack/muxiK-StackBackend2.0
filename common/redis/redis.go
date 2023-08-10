package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	fmt.Println(key, value)

	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

// Get 获取 key 对应的 value
func (rds RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			return ""
		}
	}
	return result
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds RedisClient) Del(keys ...string) error {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		return err
	}
	return nil
}
