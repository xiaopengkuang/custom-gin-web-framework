package util

import (
	"fmt"
	"github.com/xiaopengkuang/gin-web/common/web"
	"github.com/gin-gonic/gin"
)

// 构造request
func BuildRequest(gc *gin.Context) (web.Request, error) {
	// 获取url三要素
	moduleName := gc.Param("module")
	serviceName := gc.Param("service")
	operationName := gc.Param("operation")
	DoAllArgsHasText(moduleName, serviceName, operationName)

	// 获取请求参数


	// build request
	request := web.Request{Module: moduleName, Service: serviceName, Operation: operationName}
	fmt.Println(request)

	return request, nil
}
