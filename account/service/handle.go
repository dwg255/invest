package service

import (
	"game/api/thrift/gen-go/rpc"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"github.com/astaxie/beego/logs"
	"game/account/common"
	"fmt"
	"context"
	"strings"
)

func (p *UserServer) CreateNewUser(ctx context.Context, nickName string, avatarAuto string, gold int64) (r *rpc.Result_, err error) {
	logs.Debug("CreateNewUser nickName: %v", nickName)
	r = rpc.NewResult_()
	conn := userServiceConf.RedisConf.RedisPool.Get()
	defer conn.Close()
	userId, err := redis.Int(conn.Do("incr", "user_id_incr"))
	if err != nil {
		logs.Error("incr err:%v", err)
		return
	}
	logs.Debug("new user_id:%v", userId)
	if err != nil {
		r.Code = rpc.ErrorCode_UnknowError
		return
	}
	userRedisKey := userServiceConf.UserRedisPrefix + strconv.Itoa(userId)
	conn.Do("hmset", userRedisKey, "user_id", userId, "nick_name", nickName, "avatar_auto", avatarAuto, "gold", gold)
	r.UserObj = rpc.NewUserInfo()
	r.UserObj.UserId = int64(userId)
	r.UserObj.NickName = nickName
	r.UserObj.AvatarAuto = avatarAuto
	r.UserObj.Gold = gold
	r.UserObj.Token, err = common.CreateToken(r.UserObj.UserId)
	return
}

func (p *UserServer) GetUserInfoById(ctx context.Context, userId int32) (r *rpc.Result_, err error) {
	r = rpc.NewResult_()
	r.UserObj = rpc.NewUserInfo()
	logs.Debug("GetUserInfoById userId: %v", userId)
	conn := userServiceConf.RedisConf.RedisPool.Get()
	defer conn.Close()
	if err != nil {
		r.Code = rpc.ErrorCode_UnknowError
		return
	}

	userRedisKey := userServiceConf.UserRedisPrefix + strconv.Itoa(int(userId))
	userExists, err := redis.Bool(conn.Do("exists", userRedisKey))
	if err != nil {
		logs.Error("check user [%d] exists err:%v", userId, err)
		return
	}
	if !userExists {
		r.Code = rpc.ErrorCode_UserIsNull
		return
	}
	resp, err := redis.StringMap(conn.Do("hgetall", userRedisKey))
	if userIdStr, ok := resp["user_id"]; ok {
		var userId int
		userId, err = strconv.Atoi(userIdStr)
		if err != nil {
			r.Code = rpc.ErrorCode_UnknowError
			return
		}
		r.UserObj.UserId = int64(userId)
	}
	if goldStr, ok := resp["gold"]; ok {
		var gold int
		gold, err = strconv.Atoi(goldStr)
		if err != nil {
			r.Code = rpc.ErrorCode_UnknowError
			return
		}
		r.UserObj.Gold = int64(gold)
	}
	if nickName, ok := resp["nick_name"]; ok {
		r.Code = rpc.ErrorCode_UnknowError
		r.UserObj.NickName = nickName
	}
	if avatarAuto, ok := resp["avatar_auto"]; ok {
		r.UserObj.AvatarAuto = avatarAuto
	}
	r.Code = rpc.ErrorCode_Success
	logs.Debug("user obj :%v",r.UserObj)
	return
}

func (p *UserServer) GetUserInfoByken(ctx context.Context, token string) (r *rpc.Result_, err error) {
	token = strings.Replace(token, " ", "+", -1)
	logs.Debug("GetUserInfoByken token %v", token)
	userId, err := common.CheckToken(token)
	if err != nil {
		return
	}
	r, err = p.GetUserInfoById(ctx, int32(userId))
	return
}

func (p *UserServer) ModifyGoldById(ctx context.Context, behavior string, userId int32, gold int64) (r *rpc.Result_, err error) {
	logs.Debug("Modify Gold [%d] By Id [%d] behavior: %v ,", gold, userId, behavior)
	r = rpc.NewResult_()
	conn := userServiceConf.RedisConf.RedisPool.Get()
	defer conn.Close()
	userRedisKey := userServiceConf.UserRedisPrefix + strconv.Itoa(int(userId))
	userExists, err := redis.Bool(conn.Do("exists", userRedisKey))
	if err != nil {
		logs.Error("check user [%d] exists err:%v", userId, err)
		return
	}
	if !userExists {
		r.Code = rpc.ErrorCode_UserIsNull
		return
	}
	leftGold, err := redis.Int(conn.Do("hincrby", userRedisKey, "gold", gold))
	if err != nil {
		r.Code = rpc.ErrorCode_UnknowError
		logs.Error("ModifyGoldById err:%v", err)
		return
	}
	if leftGold < 0 {
		r.Code = rpc.ErrorCode_GoldNotEnough
		conn.Do("hincrby", userServiceConf.UserRedisPrefix+strconv.Itoa(int(userId)), "gold", -1*gold)
		err = fmt.Errorf("user gold not enough")
		return
	}
	r, err = p.GetUserInfoById(ctx, userId)
	return
}

func (p *UserServer) ModifyGoldByToken(ctx context.Context, behavior string, token string, gold int64) (r *rpc.Result_, err error) {
	//logs.Debug("ModifyGoldByToken token: %v", token)
	r = rpc.NewResult_()
	userId, err := common.CheckToken(token)
	if err != nil {
		r.Code = rpc.ErrorCode_VerifyError
		logs.Error("check token [%s] failed", token)
		return
	}
	r, err = p.ModifyGoldById(ctx, behavior, int32(userId), gold)
	return
}
