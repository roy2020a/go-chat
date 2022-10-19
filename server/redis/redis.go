package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

func InitRedisPool(maxIdle, maxActive int, idleTimeout time.Duration, host string) *redis.Pool {
	pool := &redis.Pool{
		// 初始化链接数量
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host)
		},
	}
	return pool

}
