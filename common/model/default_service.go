package model

import (
	"fmt"
	"gin-web/common/web"
	"reflect"
	"strings"
	"github.com/gin-gonic/gin"
	"gin-web/util"
)

const (
	APIPrefix = "API_"
)

// TODO: looking for An elegant way
var APIRequestType = []string{"web.Request"}
var APIResponseType = []string{"interface {}", "error"}

// 对外service 命名规则为API_operation
type ServiceInterface interface {
	Name() string
	// ...
}

// <methodName,method reflect value>
type Service struct {
	APIs map[string]reflect.Value
}

func (s *Service) initService() {
	// 初始化service api map
	s.APIs = make(map[string]reflect.Value)
}

func (s *Service) convertService(targetService ServiceInterface) error {
	if s.APIs == nil {
		s.initService()
	}

	serviceType := reflect.TypeOf(targetService)
	methodNum := serviceType.NumMethod()
	if methodNum == 0 {
		return nil
	}

	serviceValue := reflect.ValueOf(targetService)
	for i := 0; i < methodNum; i++ {
		methodName := serviceType.Method(i).Name
		if !strings.HasPrefix(methodName, APIPrefix) {
			continue
		}

		method := serviceValue.MethodByName(methodName)
		if !isMethodForAPI(method) {
			continue
		}

		// 到此可以决定method 是否是对外api
		s.addAPI(methodName, method)
	}

	return nil
}

func isMethodForAPI(method reflect.Value) bool {
	// 判断方法是否有效
	if method.IsNil() || !method.IsValid() {
		return false
	}

	// 判断方法的参数
	canInterface := method.CanInterface()
	if canInterface == false {
		return false
	}

	//判断方法的入参和出参
	if !isValidAPI(method) {
		return false
	}

	return true
}

func (s *Service) addAPI(methodName string, methodReflectValue reflect.Value) {
	if strings.HasPrefix(methodName, APIPrefix) {
		methodName = strings.Replace(methodName, APIPrefix, "", 1)
	}

	if s.APIs == nil {
		s.initService()
	}

	s.APIs[methodName] = methodReflectValue
}

// 判断方法是否是标准的tgrpc响应接口参数：func(tgrpcbase.Request) (tgrpcbase.Response)
func isValidAPI(methodValue reflect.Value) bool {
	targetInterface := methodValue.Interface()

	if targetInterface == nil {
		return false
	}

	targetType := reflect.TypeOf(targetInterface)
	targetKind := targetType.Kind()

	if targetKind != reflect.Func {
		return false
	}

	if targetType.NumIn() != len(APIRequestType) || targetType.NumOut() != len(APIResponseType) {
		return false
	}

	// 验证request
	for i := 0; i < len(APIRequestType) && i < targetType.NumIn(); i++ {
		if !strings.EqualFold(APIRequestType[i], targetType.In(i).String()) {
			return false
		}
	}

	// 验证response
	for i := 0; i < len(APIResponseType) && i < targetType.NumOut(); i++ {
		if !strings.EqualFold(APIResponseType[i], targetType.Out(i).String()) {
			return false
		}
	}

	return true
}

func (s *Service) Execute(gc *gin.Context) web.Response {
	// 构建web request
	request, err := util.BuildRequest(gc)
	if err != nil {
		return web.FailResp(err.Error(), nil)
	}

	if s.APIs == nil {
		return web.DenyResp("No method in the service : "+request.Service, nil)
	}

	method, ok := s.APIs[request.Operation]
	if !ok {
		return web.DenyResp("no method matched", nil)
	}

	// 再次校验method
	if !isValidAPI(method) {
		return web.DenyResp("invalid method", nil)
	}

	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(request)
	funcResult := method.Call(in)

	// 0 interface, 1 error
	if funcResult == nil || len(funcResult) != 2 {
		return web.FailResp("fail to get response", nil)
	}

	data := funcResult[0].Interface()
	errorInfo := funcResult[1].Interface()

	//  与下面重复
	if errorInfo == nil {
		return web.SuccessResp(data)
	}

	switch errorInfo.(type) {
	case error:
		break
	case nil:
		break
	default:
		fmt.Println(errorInfo)
		return web.FailResp("invalid error type", nil)
	}

	if errorInfo != nil {
		// 将第二个参数转换成error
		errorMsg := errorInfo.(error)
		return web.FailResp(errorMsg.Error(), data)
	}

	return web.SuccessResp(data)
}
