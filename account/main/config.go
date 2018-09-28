package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
	"game/account/common"
)

var (
	UserServiceConf = &common.UserServiceConf{}
)

func initConf() (err error) {
	UserServiceConf.ThriftPort,err = beego.AppConfig.Int("thrift_port")
	if err != nil {
		logs.Error("init http_port failed,err:%v",err)
		return
	}
	logs.Debug("read conf succ,http port %v", UserServiceConf.ThriftPort)

	//todo redis相关配置
	UserServiceConf.RedisConf.RedisAddr = beego.AppConfig.String("redis_addr")
	if len(UserServiceConf.RedisConf.RedisAddr) == 0 {
		err = fmt.Errorf("init config failed, http_addr [%s]", UserServiceConf.RedisConf.RedisAddr)
		return
	}

	UserServiceConf.RedisConf.RedisAddr = beego.AppConfig.String("redis_addr")
	if len(UserServiceConf.RedisConf.RedisAddr) == 0 {
		err = fmt.Errorf("init config failed, redis_addr [%s]", UserServiceConf.RedisConf.RedisAddr)
		return
	}
	UserServiceConf.RedisConf.RedisMaxIdle,err = beego.AppConfig.Int("redis_max_idle")
	if err != nil {
		logs.Error("init config failed,read redis_max_idle err :%v", err)
		return
	}
	UserServiceConf.RedisConf.RedisMaxActive, err = beego.AppConfig.Int("redis_max_active")
	if err != nil {
		logs.Error("init config failed,read redis_max_active err :%v", err)
		return
	}
	UserServiceConf.RedisConf.RedisIdleTimeout, err = beego.AppConfig.Int("redis_idle_timeout")
	if err != nil {
		logs.Error("init config failed,read redis_idle_timeout err :%v", err)
		return
	}

	//todo 工程密钥
	UserServiceConf.AppSecret = beego.AppConfig.String("app_secret")
	if len(UserServiceConf.AppSecret) == 0 {
		err = fmt.Errorf("init config failed, app_secret [%s]", UserServiceConf.AppSecret)
		return
	}

	//todo 日志配置
	UserServiceConf.LogPath = beego.AppConfig.String("log_path")
	UserServiceConf.LogPath = beego.AppConfig.String("log_path")
	UserServiceConf.UserRedisPrefix = beego.AppConfig.String("user_redis_prefix")

	return
}
