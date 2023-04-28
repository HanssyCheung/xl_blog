package main

import (
	"xl-go-blog/core"
	"xl-go-blog/global"
)

func main() {
	//读取配置文件
	core.InitConf()
	global.DB = core.InitGorm()
}
