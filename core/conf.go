package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
	"xl-go-blog/config"
	"xl-go-blog/global"
)

//ConfigFile 在这个里面读取yaml文件配置，将基础的信息加载进来
const ConfigFile = "settings.yaml"

func InitConf() {
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

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
