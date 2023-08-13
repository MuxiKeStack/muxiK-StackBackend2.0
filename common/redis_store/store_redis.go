package redis_store

import (
	rds "github.com/MuxiKeStack/muxiK-StackBackend2.0/common/redis"
	"time"
)

type RedisStore struct {
	RedisClient *rds.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value interface{}, expire time.Duration) error {

	return s.RedisClient.Set(s.KeyPrefix+key, value, expire)
}

func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

// 检查验证码
func (s *RedisStore) Verify(id, answer string) bool {
	v := s.Get(id, false)
	if v == answer {
		s.Get(id, true)
	}
	return v == answer
}
