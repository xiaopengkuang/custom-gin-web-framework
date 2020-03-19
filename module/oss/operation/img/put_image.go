package img

import (
	"github.com/gin-gonic/gin"
)

const (
	ImageServicePutOperation = "put"
)

type PutImg struct {
}

func (p *PutImg) OperationName() string {
	return ImageServicePutOperation
}

func (p *PutImg) Process(ctx *gin.Context) (interface{}, error) {

	return "PutFile", nil
}
