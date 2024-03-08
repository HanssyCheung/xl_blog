package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"xl-go-blog/core"
	"xl-go-blog/global"
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
	global.Log.Error("哈哈哈")
	global.Log.Infof("呵呵呵")
	//测试日志
	logrus.Warnln("嘻嘻嘻")
	logrus.Error("哈哈哈")
	logrus.Infof("呵呵呵")
}
