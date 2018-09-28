package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
	"encoding/json"
	"game/account/service"
)

//var (
//	redisPool *redis.Pool
//	etcdClient *etcd_client.Client
//)

func initRedis() (err error) {
	UserServiceConf.RedisConf.RedisPool = &redis.Pool{
		MaxIdle: UserServiceConf.RedisConf.RedisMaxIdle,
		MaxActive:UserServiceConf.RedisConf.RedisMaxActive,
		IdleTimeout:time.Duration(UserServiceConf.RedisConf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn ,error) {
			return redis.Dial("tcp",UserServiceConf.RedisConf.RedisAddr)
		},
	}
	conn := UserServiceConf.RedisConf.RedisPool.Get()
	defer conn.Close()
	//defer UserServiceConf.RedisConf.RedisPool.Close()
	_,err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed,err:%v",err)
		return
	}
	return
}

func converLogLevel(logLevel string) int {
	switch logLevel {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}
func initLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = UserServiceConf.LogPath
	config["level"] = converLogLevel(UserServiceConf.LogLevel)

	configStr,err := json.Marshal(config)
	if err != nil {
		fmt.Println("marsha1 faild,err",err)
		return
	}
	logs.SetLogger(logs.AdapterFile,string(configStr))
	return
}


func initSec() (err error) {
	err = initLogger()
	if err != nil {
		logs.Error("init logger failed,err:%v",err)
		return
	}
	err = initRedis()
	if err != nil {
		logs.Error("init redis failed,err :%v",err)
		return
	}

	service.InitService(UserServiceConf)
	logs.Info("init sec succ")
	return
}

