package engine

import (
	"fmt"
	"gin-web/common/model"
	"gin-web/common/web"
	"log"
	"github.com/gin-gonic/gin"
)

// module register error
var registerError []error

// module factory
var moduleFactory map[string]*model.Module

// 初始化
func init() {
	initModuleFactory()
	initRegisterError()
}

// 初始化module
func initModuleFactory() {
	moduleFactory = make(map[string]*model.Module)
}

// 初始化RegisterError
func initRegisterError() {
	registerError = make([]error, 0)
}

// 添加注册异常
func addRegisterError(err error) {
	if registerError == nil {
		initRegisterError()
	}

	registerError = append(registerError, err)
}

// 添加注册异常
func addRegisterErrors(errs []error) {
	if registerError == nil {
		initRegisterError()
	}

	if errs == nil || len(errs) == 0 {
		return
	}

	for i := 0; i < len(errs); i++ {
		registerError = append(registerError, errs[i])
	}
}

// 获取module
func GetModule(name string) (*model.Module, bool) {
	if moduleFactory == nil {
		return nil, false
	}

	module, ok := moduleFactory[name]

	return module, ok
}

// 注册module
func RegisterModule(name string, module *model.Module) {
	if moduleFactory == nil {
		initModuleFactory()
	}

	newModuleErrs := module.GetRegisterError()
	if newModuleErrs != nil && len(newModuleErrs) > 0 {
		addRegisterErrors(newModuleErrs)
		return
	}

	_, ok := moduleFactory[name]
	if ok {
		err := fmt.Errorf(name, " already exist")
		log.Println(err)
		addRegisterError(err)
		return
	}

	moduleFactory[name] = module
}

// 调用方法
func ExecuteFunction(gc *gin.Context) web.Response {
	moduleName := gc.Param("module")
	serviceName := gc.Param("service")

	// 校验module
	module, ok := GetModule(moduleName)
	if !ok || module == nil {
		return web.DenyResp("module not matched", nil)
	}

	// 校验service
	service, err := module.Get(serviceName)
	if err != nil {
		log.Println(err)
		return web.DenyResp("service not match", nil)
	}

	// 执行响应方法方法
	response := service.Execute(gc)

	return response
}

func GetRegisterError() []error {
	return registerError
}
