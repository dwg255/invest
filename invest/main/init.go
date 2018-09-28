package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"time"
	etcd_client "github.com/coreos/etcd/clientv3"
	"fmt"
	"encoding/json"
	"game/invest/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

func initRedis() (err error) {
	gameConf.RedisConf.RedisPool = &redis.Pool{
		MaxIdle: gameConf.RedisConf.RedisMaxIdle,
		MaxActive:gameConf.RedisConf.RedisMaxActive,
		IdleTimeout:time.Duration(gameConf.RedisConf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn ,error) {
			return redis.Dial("tcp",gameConf.RedisConf.RedisAddr)
		},
	}
	conn := gameConf.RedisConf.RedisPool.Get()
	defer conn.Close()
	_,err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed,err:%v",err)
		return
	}
	return
}

func initEtcd () (err error) {
	cli,err := etcd_client.New(etcd_client.Config{
		Endpoints:[]string{gameConf.EtcdConf.EtcdAddr},
		DialTimeout:time.Duration(gameConf.EtcdConf.Timeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect Etcd failed,err :",err)
		return
	}
	gameConf.EtcdConf.EtcdClient = cli
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
	config["filename"] = gameConf.LogPath
	config["level"] = converLogLevel(gameConf.LogLevel)

	configStr,err := json.Marshal(config)
	if err != nil {
		fmt.Println("marsha1 faild,err",err)
		return
	}
	logs.SetLogger(logs.AdapterFile,string(configStr))
	return
}

func initMysql() (err error) {
	conf := gameConf.MysqlConf
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",conf.MysqlUser,conf.MysqlPassword,conf.MysqlAddr,conf.MysqlDatabase)
	logs.Debug(dsn)
	database, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return
	}

	gameConf.MysqlConf.Pool = database
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
	/*err = initEtcd()
	if err != nil {
		logs.Error("init etcd failed,err:%v",err)
		return
	}*/
	err = initMysql()
	if err != nil {
		logs.Error("init mysql failed,err :%v",err)
		return
	}
	service.InitService(gameConf)
	logs.Info("init sec succ")
	return
}

