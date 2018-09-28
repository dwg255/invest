package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
	"game/invest/common"
)

var (
	gameConf = &common.GameConf{}
)

func initConf() (err error) {
	gameConf.HttpPort,err = beego.AppConfig.Int("http_port")
	if err != nil {
		logs.Error("init http_port failed,err:%v",err)
		return
	}
	logs.Debug("read conf succ,http port %v", gameConf.HttpPort)

	// redis相关配置
	gameConf.RedisConf.RedisAddr = beego.AppConfig.String("redis_addr")
	if len(gameConf.RedisConf.RedisAddr) == 0 {
		err = fmt.Errorf("init config failed, redis_addr [%s]", gameConf.RedisConf.RedisAddr)
		return
	}
	gameConf.RedisConf.RedisMaxIdle,err = beego.AppConfig.Int("redis_max_idle")
	if err != nil {
		logs.Error("init config failed,read redis_max_idle err :%v", err)
		return
	}
	gameConf.RedisConf.RedisMaxActive, err = beego.AppConfig.Int("redis_max_active")
	if err != nil {
		logs.Error("init config failed,read redis_max_active err :%v", err)
		return
	}
	gameConf.RedisConf.RedisIdleTimeout, err = beego.AppConfig.Int("redis_idle_timeout")
	if err != nil {
		logs.Error("init config failed,read redis_idle_timeout err :%v", err)
		return
	}
	gameConf.RedisKey.RedisKeyUserStake = beego.AppConfig.String("redis_key_user_stake_msg")
	if len(gameConf.RedisKey.RedisKeyUserStake) == 0 {
		err = fmt.Errorf("init config failed, redis_key_user_stake_msg [%s]", gameConf.RedisKey.RedisKeyUserStake)
		return
	}
	gameConf.RedisKey.RedisKeyInvestBase = beego.AppConfig.String("redis_key_invest_base")
	if len(gameConf.RedisKey.RedisKeyInvestBase) == 0 {
		err = fmt.Errorf("init config failed, redis_key_invest_base [%s]", gameConf.RedisKey.RedisKeyInvestBase)
		return
	}
	// mysql相关配置
	gameConf.MysqlConf.MysqlAddr = beego.AppConfig.String("mysql_addr")
	if len(gameConf.MysqlConf.MysqlAddr) == 0 {
		err = fmt.Errorf("init config failed, mysql_addr [%s]", gameConf.MysqlConf.MysqlAddr)
		return
	}
	gameConf.MysqlConf.MysqlUser = beego.AppConfig.String("mysql_user")
	if len(gameConf.MysqlConf.MysqlUser) == 0 {
		err = fmt.Errorf("init config failed, mysql_user [%s]", gameConf.MysqlConf.MysqlUser)
		return
	}
	gameConf.MysqlConf.MysqlPassword = beego.AppConfig.String("mysql_password")
	if len(gameConf.MysqlConf.MysqlPassword) == 0 {
		err = fmt.Errorf("init config failed, mysql_password [%s]", gameConf.MysqlConf.MysqlPassword)
		return
	}
	gameConf.MysqlConf.MysqlDatabase = beego.AppConfig.String("mysql_db")
	if len(gameConf.MysqlConf.MysqlDatabase) == 0 {
		err = fmt.Errorf("init config failed, mysql_password [%s]", gameConf.MysqlConf.MysqlDatabase)
		return
	}

	// etcd相关配置
	gameConf.EtcdConf.EtcdAddr = beego.AppConfig.String("etcd_addr")
	if len(gameConf.EtcdConf.EtcdAddr) == 0 {
		err = fmt.Errorf("init config failed, etcd_addr [%s]", gameConf.EtcdConf.EtcdAddr)
		return
	}
	gameConf.EtcdConf.Timeout, err = beego.AppConfig.Int("etcd_timeout")
	if err != nil {
		err = fmt.Errorf("init config failed,read etcd_timeout err :%v", err)
		return
	}

	// 密钥
	gameConf.AppSecret = beego.AppConfig.String("app_secret")
	if len(gameConf.AppSecret) == 0 {
		err = fmt.Errorf("init config failed, app_secret [%s]", gameConf.AppSecret)
		return
	}

	// 日志配置
	gameConf.LogPath = beego.AppConfig.String("log_path")
	gameConf.LogLevel = beego.AppConfig.String("log_level")

	//todo 配置到etcd
	gameConf.PumpingRate,err = beego.AppConfig.Int("pumping_rate")
	if err != nil {
		err = fmt.Errorf("init config failed,read pumping_rate err :%v", err)
		return
	}
	if gameConf.PumpingRate < 0 || gameConf.PumpingRate > 10 {
		err = fmt.Errorf("init config failed,config pumping_rate [%d] out of range ", gameConf.PumpingRate)
		return
	}
	return
}
