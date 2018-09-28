package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
	"game/invest/common"
)

var (
	backConf = &common.GameConf{}
)

func initConf() (err error) {
	//todo redis相关配置
	backConf.RedisConf.RedisAddr = beego.AppConfig.String("redis_addr")
	if len(backConf.RedisConf.RedisAddr) == 0 {
		err = fmt.Errorf("init config failed, redis_addr [%s]", backConf.RedisConf.RedisAddr)
		return
	}
	backConf.RedisConf.RedisMaxIdle,err = beego.AppConfig.Int("redis_max_idle")
	if err != nil {
		logs.Error("init config failed,read redis_max_idle err :%v", err)
		return
	}
	backConf.RedisConf.RedisMaxActive, err = beego.AppConfig.Int("redis_max_active")
	if err != nil {
		logs.Error("init config failed,read redis_max_active err :%v", err)
		return
	}
	backConf.RedisConf.RedisIdleTimeout, err = beego.AppConfig.Int("redis_idle_timeout")
	if err != nil {
		logs.Error("init config failed,read redis_idle_timeout err :%v", err)
		return
	}
	backConf.RedisKey.RedisKeyUserStake = beego.AppConfig.String("redis_key_user_stake_msg")
	if len(backConf.RedisKey.RedisKeyUserStake) == 0 {
		err = fmt.Errorf("init config failed, redis_key_user_stake_msg [%s]", backConf.RedisKey.RedisKeyUserStake)
		return
	}
	backConf.RedisKey.RedisKeyInvestBase = beego.AppConfig.String("redis_key_invest_base")
	if len(backConf.RedisKey.RedisKeyInvestBase) == 0 {
		err = fmt.Errorf("init config failed, redis_key_invest_base [%s]", backConf.RedisKey.RedisKeyInvestBase)
		return
	}
	//todo mysql相关配置
	backConf.MysqlConf.MysqlAddr = beego.AppConfig.String("mysql_addr")
	if len(backConf.MysqlConf.MysqlAddr) == 0 {
		err = fmt.Errorf("init config failed, mysql_addr [%s]", backConf.MysqlConf.MysqlAddr)
		return
	}
	backConf.MysqlConf.MysqlUser = beego.AppConfig.String("mysql_user")
	if len(backConf.MysqlConf.MysqlUser) == 0 {
		err = fmt.Errorf("init config failed, mysql_user [%s]", backConf.MysqlConf.MysqlUser)
		return
	}
	backConf.MysqlConf.MysqlPassword = beego.AppConfig.String("mysql_password")
	if len(backConf.MysqlConf.MysqlPassword) == 0 {
		err = fmt.Errorf("init config failed, mysql_password [%s]", backConf.MysqlConf.MysqlPassword)
		return
	}
	backConf.MysqlConf.MysqlDatabase = beego.AppConfig.String("mysql_db")
	if len(backConf.MysqlConf.MysqlDatabase) == 0 {
		err = fmt.Errorf("init config failed, mysql_password [%s]", backConf.MysqlConf.MysqlDatabase)
		return
	}
	//todo 日志配置
	backConf.LogPath = beego.AppConfig.String("log_path")
	backConf.LogLevel = beego.AppConfig.String("log_level")

	return
}
