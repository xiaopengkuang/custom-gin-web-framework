package config

import (
	"gin-web/config/model"
	"gin-web/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var AppConfig *model.AppConfig

// 初始化配置
func init() {
	loadAppConfig()
}

func loadAppConfig() {
	// 获取gopath
	gopathSrc, err := util.GetGoPathSrc()
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	// build配置文件路径
	configFilePath := gopathSrc + appConfigRelativePath + configFileName
	err = util.PathExist(configFilePath)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	// 读取文件
	configData, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	AppConfig = &model.AppConfig{}

	err = yaml.Unmarshal([]byte(configData), &AppConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}
}

// 设置日志配置
func initLog() {
	//log.SetOutput(&lumberjack.Logger{
	//	Filename:   "/Users/kuangxiaopeng/Workspace/log/go/gin-web/foo.log",
	//	MaxSize:    500, // megabytes
	//	MaxBackups: 3,
	//	MaxAge:     28,   //days
	//	Compress:   true, // disabled by default
	//})
}
