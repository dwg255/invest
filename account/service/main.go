package service

import "game/account/common"
var(
	userServiceConf *common.UserServiceConf
)
func InitService(UserServiceConf *common.UserServiceConf)  {
	userServiceConf = UserServiceConf
}
type UserServer struct {

}