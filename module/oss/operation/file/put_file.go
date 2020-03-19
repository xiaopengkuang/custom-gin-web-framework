package file

import (
	"github.com/gin-gonic/gin"
)

const (
	FileServicePutOperation = "put"
)

type PutFile struct {
}

func (p *PutFile) OperationName() string {
	return FileServicePutOperation
}

func (p *PutFile) Process(ctx *gin.Context) (interface{}, error) {

	return "PutFile", nil
}
