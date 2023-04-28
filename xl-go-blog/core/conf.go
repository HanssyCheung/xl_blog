package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"xl-go-blog/config"
	"xl-go-blog/global"
)

//在这个里面读取yaml文件配置

func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c) //解析
	if err != nil {
		log.Fatal("config Init Unmarshal: %V", err)
	}
	log.Println("config yamlFile load Init success.")
	fmt.Println(c)
	global.Config = c
}
