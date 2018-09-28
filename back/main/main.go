package main

import (
	"github.com/astaxie/beego/logs"
)

func main()  {
	err := initConf()
	if err != nil {
		logs.Error("init conf err:%v",err)
		return
	}

	err = initSec()
	if err != nil {
		logs.Error("init sec err:%v",err)
		return
	}
}