package model

import (
	"fmt"
	"gin-web/common/web"
	"github.com/gin-gonic/gin"
)

type Module struct {
	name          string
	ServiceMap    map[string]ModuleService
	registerError []error
}

type ModuleService interface {
	Execute(gc *gin.Context) web.Response
}

//  注册服务
func (m *Module) Register(targetService ServiceInterface) {
	if m.ServiceMap == nil {
		m.initServiceMap()
	}

	if targetService == nil {
		m.addRegisterError(fmt.Errorf("service not init!\n"))
		return
	}

	// 获取service name
	serviceName := targetService.Name()

	_, ok := m.ServiceMap[serviceName]
	if ok {
		m.addRegisterError(fmt.Errorf("module already exist"))
		return
	}

	service := &Service{}
	err := service.convertService(targetService)
	if err != nil {
		m.addRegisterError(err)
		return
	}

	// 添加service到map
	m.ServiceMap[serviceName] = service
}

func (m *Module) RegisterGinService(serviceName string, apis map[string]GinOperation) {
	// check
	m.checkServiceMap()
	_, ok := m.ServiceMap[serviceName]
	if ok {
		m.addRegisterError(fmt.Errorf("module already exist"))
		return
	}

	service := &GinService{ServiceName: serviceName, APIs: apis}
	m.ServiceMap[serviceName] = service
}

func (m *Module) Get(serviceName string) (ModuleService, error) {
	if m.ServiceMap == nil {
		return nil, fmt.Errorf("service map not initServiceMap")
	}

	service, ok := m.ServiceMap[serviceName]
	if !ok || service == nil {
		return nil, fmt.Errorf("service not exist")
	}

	return service, nil
}

func (m *Module) addRegisterError(err error) {
	if m.registerError == nil {
		m.registerError = make([]error, 0)
	}

	m.registerError = append(m.registerError, err)
}

func (m *Module) initServiceMap() {
	m.ServiceMap = make(map[string]ModuleService)
	m.registerError = make([]error, 0)
}

func (m *Module) GetRegisterError() []error {
	return m.registerError
}

func (m *Module) checkServiceMap() {
	if m.ServiceMap == nil {
		m.ServiceMap = make(map[string]ModuleService)
	}

	if m.registerError == nil {
		m.registerError = make([]error, 0)
	}
}
