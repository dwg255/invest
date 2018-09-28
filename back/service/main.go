package service

import (
	"game/invest/common"
	"sync"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"time"
)

var (
	backConf *common.GameConf
	wg sync.WaitGroup
)
func Run(conf *common.GameConf)  {
	backConf = conf
	wg.Add(1)
	go InvestUserStake()
	wg.Add(1)
	go InvestBase()
	wg.Wait()
}

func InvestUserStake()  {
	var conn redis.Conn
	conn = backConf.RedisConf.RedisPool.Get()
	for {
		resp,err := redis.String(conn.Do("RPOP",backConf.RedisKey.RedisKeyUserStake))
		if err != nil {
			time.Sleep(time.Second)
			logs.Error("get invest user stake message failed ,rpop [%v] err:%v",backConf.RedisKey.RedisKeyUserStake,err)
			continue
		}
		investUserStake := &common.InvestUserStake{}
		err = json.Unmarshal([]byte(resp),investUserStake)
		if err != nil {
			logs.Error("unmarsha1 invest user stake message [%v] failed err:%v",resp,err)
			continue
		}
		_,err = backConf.MysqlConf.Pool.Exec(`insert into game_invest_user_stake(
game_times_id,periods,room_id,room_type,user_id,nickname,user_all_stake,get_gold,stake_detail,game_result,game_pool,last_stake_time) 
values(?,?,?,?,?,?,?,?,?,?,?,?) `,investUserStake.GameTimesId,
			investUserStake.Periods,
			investUserStake.RoomId,
			investUserStake.RoomType,
			investUserStake.UserId,
			investUserStake.Nickname,
			investUserStake.UserAllStake,
			investUserStake.WinGold,
			investUserStake.StakeDetail,
			investUserStake.GameResult,
			investUserStake.Pool,
			investUserStake.StakeTime,
			)
		if err != nil {
			logs.Error("insert user stake [%v],err :%v",investUserStake,err)
			continue
		}
	}
	conn.Close()
	wg.Done()
}
func InvestBase()  {
	var conn redis.Conn
	conn = backConf.RedisConf.RedisPool.Get()
	for {
		resp,err := redis.String(conn.Do("RPOP",backConf.RedisKey.RedisKeyInvestBase))
		if err != nil {
			//logs.Debug("%T",backConf.RedisKey.RedisKeyInvestBase)
			time.Sleep(time.Second)
			logs.Error("get invest base failed ,rpop [%v] resp[%v] err:%v",backConf.RedisKey.RedisKeyInvestBase,resp,err)
			continue
		}
		investBase := &common.InvestBase{}
		err = json.Unmarshal([]byte(resp),investBase)
		if err != nil {
			logs.Error("unmarsha1 invest user stake message [%v] failed err:%v",resp,err)
			continue
		}
		_,err = backConf.MysqlConf.Pool.Exec(`insert into game_invest_base(
game_times_id,periods,game_pool,stake_detail,game_result,start_time) 
values(?,?,?,?,?,?) `,investBase.GameTimesId,
			investBase.Periods,
			investBase.Pool,
			investBase.StakeDetail,
			investBase.GameResult,
			investBase.StartTime,
		)
		if err != nil {
			logs.Error("insert invest base [%v],err :%v",investBase,err)
			continue
		}
	}
	conn.Close()
	wg.Done()
}