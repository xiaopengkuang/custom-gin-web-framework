package oss

import (
	"github.com/xiaopengkuang/gin-web/common/model"
	"github.com/xiaopengkuang/gin-web/module/oss/service"
)

const (
	OSSModuleName = "oss"
)

var userModule *model.Module

func init() {
	registerService()
}

// 注册service
func registerService() {
	userModule = &model.Module{}
	userModule.RegisterGinService(service.ImageServiceName, service.GetFileService().APIs)
	userModule.RegisterGinService(service.FileServiceName, service.GetImageService().APIs)
}

func GetModule() *model.Module {
	return userModule
}
