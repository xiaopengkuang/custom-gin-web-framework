package base

import (
	"gin-web/common/model"
	"gin-web/module/base/service"
)

const (
	BaseModuleName = "base"
)

var userModule *model.Module

func init() {
	registerService()
}

// 注册service
func registerService() {
	userModule = &model.Module{}
	userModule.Register(&service.UserService{})
}

func GetModule() *model.Module {
	return userModule
}
