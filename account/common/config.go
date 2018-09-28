package common

import "github.com/garyburd/redigo/redis"


type UserServiceConf struct {
	ThriftPort  int
	RedisConf RedisConf
	LogPath   string
	LogLevel  string
	AppSecret string
	UserRedisPrefix string
}

type RedisConf struct {
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
	RedisPool        *redis.Pool
}
