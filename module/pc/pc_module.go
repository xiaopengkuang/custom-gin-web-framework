package pc

import (
	"gin-web/common/model"
	"gin-web/module/pc/service"
)

const (
	PCModuleName = "pc"
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
