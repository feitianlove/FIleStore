package pool

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	pool      *redis.Pool
	redisHost = "127.0.0.1:6379"
	redisPass = ""
)

func NewRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:         50,
		MaxActive:       30,
		IdleTimeout:     300 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				panic(err)
				return nil, err
			}
			if _, err := c.Do("AUTH", redisPass); err != nil {
				return nil, err
			}
			return c, nil
		},
		// 每分钟去检测redis 链接池是否可用
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			if _, err := c.Do("PING"); err != nil {
				return err
			}
			return nil
		},
	}

}

func init() {
	pool = NewRedisPool()
}
func RedisPool() *redis.Pool {
	return pool
}
