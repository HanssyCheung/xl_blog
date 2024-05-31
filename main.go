package main

import (
	"fmt"
	"xl-go-blog/core"
	"xl-go-blog/flag"
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

	//命令行参数绑定
	option := flag.Parse()
	fmt.Println(option)
	if flag.IsWebStop(option) {
		fmt.Println("命令行参数打印")
		flag.SwitchOption(option)
		return
	}

	//路由
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("xl-server运行在：%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatal(err.Error())
	}
}
