package model

import (
	"gin-web/common/web"
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
)

type GinService struct {
	ServiceName           string
	APIs                  map[string]GinOperation
	registerOperationErrs []error
}

func (g *GinService) Name() string {
	return g.ServiceName
}

func (g *GinService) Execute(gc *gin.Context) web.Response {
	// 获取响应的operation
	operationName := gc.Param("operation")
	operation, err := g.getOperation(operationName)
	if err != nil {
		log.Println(err.Error())
		web.FailResp(err.Error(), nil)
	}

	result, err := operation.Process(gc)
	if err != nil {
		log.Println(err.Error())
		web.FailResp(err.Error(), nil)
	}

	return web.SuccessResp(result)
}

func (g *GinService) getOperation(name string) (GinOperation, error) {
	if g.APIs == nil {
		return nil, fmt.Errorf("No API to match\n")
	}

	operation, ok := g.APIs[name]
	if !ok {
		return nil, fmt.Errorf("method not matched")
	}

	return operation, nil
}
