package lib

import (
	"github.com/go-redis/redis"
	"strconv"
)

type RedisConnector struct {
	Host string
	Port int
}

func (c *RedisConnector) Connect() *RedisClient {
	return &RedisClient{
		redis.NewClient(&redis.Options{
			Addr: c.Host + ":" + strconv.Itoa(c.Port),
		}),
	}
}

type RedisClient struct {
	*redis.Client
}
