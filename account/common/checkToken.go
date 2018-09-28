package common

import (
	"strings"
	"strconv"
	"github.com/astaxie/beego/logs"
	"game/invest/tools"
	"fmt"
	"time"
	"net/url"
)

var(
	AesTool *tools.AesEncrypt
)
func init()  {
	var err error
	AesTool,err = tools.NewAesTool("a26fe778a427b45901dc2a9f198c4f57")
	if err != nil {
		msg := fmt.Sprintf("init aestool failed,err:%v",err)
		panic(msg)
	}
}
func CheckToken(token string) (userId int,err error) {
	token, _=url.QueryUnescape(token)
	token = strings.Replace(token," ","+",-1)
	identity,err := AesTool.Decrypt(token)
	if err != nil {
		logs.Error("token [%s] decrypt err:%v",token,err)
		return
	}
	//logs.Debug("decrypt identity :%s",identity)

	data := strings.Split(identity,"|")
	if len(data) == 2 {
		userId,err = strconv.Atoi(data[0])
		if err != nil {
			logs.Error("userid [%s] sting to int err:%v",data[0],err)
			return
		}
		return
	} else {
		err = fmt.Errorf("check token [%s] failed",token)
	}
	return
}

func CreateToken(userId int64) (token string,err error) {
	identity := fmt.Sprintf("%d|%d",userId,time.Now().Unix())
	token,err = AesTool.Encrypt(identity)
	return
}