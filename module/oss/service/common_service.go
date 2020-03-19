package service

import (
	"gin-web/common/model"
	"fmt"
)

const (
	CommonServiceName = "common"
)


type CommonService struct {
	ServiceName           string
	APIs                  map[string]model.GinOperation
	registerOperationErrs []error
}

func (f *CommonService) Name() string {
	return FileServiceName
}


func (g *CommonService) RegisterOperation(operation model.GinOperation) {
	g.checkService()
	operationName := operation.OperationName()
	_, ok := g.APIs[operationName]
	if ok {
		g.addRegisterError(fmt.Errorf("operation %s is already exist", operationName))
		return
	}

	g.APIs[operationName] = operation
}

func (g *CommonService) checkService() {
	if g.APIs == nil {
		g.APIs = make(map[string]model.GinOperation)
	}

	if g.registerOperationErrs == nil {
		g.registerOperationErrs = make([]error, 0)
	}
}

func (g *CommonService) addRegisterError(err error) {
	if g.registerOperationErrs == nil {
		g.registerOperationErrs = make([]error, 0)
	}

	g.registerOperationErrs = append(g.registerOperationErrs, err)
}
