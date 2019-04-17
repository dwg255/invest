# game
golang开发的h5小游戏 [服务端] domo

#### 软件架构
account 模块提供用户注册、验证、扣费thrift rpc服务

invest 模块处理前台http、websocket请求，启动定时器更替游戏状态。前台hub层处理用户登录退出，room层处理房间消息推送、结算。watch etcd进行服务器降级（暂无）

back 模块将数据持久化到mysql

#### 安装方式
1. 切换到在GOPATH目录下
2. git clone https://github.com/dwg255/invest.git
3. 编译帐号服务 go build -o game/account/account game/account/main
4. 编译逻辑服务 go build -o game/invest/invest game/invest/main
5. 编译持久化服务 go build -o game/back/back game/back/main
6. 启动全部
7. 简易调试demo ws.html

#### 示例链接 
http://invest.blzz.shop/login/

![](https://raw.githubusercontent.com/dwg255/invest/master/1.png)
![](https://raw.githubusercontent.com/dwg255/invest/master/1.png)
![](https://raw.githubusercontent.com/dwg255/invest/master/3.png)
