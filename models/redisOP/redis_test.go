package redisOP

import "testing"

func TestRedisLimit(t *testing.T) {
	RedisSetExpireLimit("test", 60, 10)
}
