package main

import (
	"fmt"
	"xl-go-blog/core"
	"xl-go-blog/global"
	"xl-go-blog/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
	//连接数据库
	global.DB = core.InitGorm()

	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("嘻嘻嘻")

	//路由
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("xl-server运行在：%s", addr)
	router.Run(addr)
}
