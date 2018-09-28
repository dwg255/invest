package main

import (
	"game/api/thrift/gen-go/rpc"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/astaxie/beego/logs"
	"fmt"
	"game/account/service"
)

func main() {
	err := initConf()
	if err != nil {
		logs.Error("init conf err:%v",err)
		return
	}
	err = initSec()
	if err != nil {
		logs.Error("initSec err:%v",err)
		return
	}

	port := fmt.Sprintf(":%d",UserServiceConf.ThriftPort)
	transport, err := thrift.NewTServerSocket(port)
	if err != nil {
		panic(err)
	}
	handler := &service.UserServer{}
	processor := rpc.NewUserServiceProcessor(handler)

	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
