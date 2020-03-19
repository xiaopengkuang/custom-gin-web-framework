package model

import "github.com/gin-gonic/gin"

type GinOperation interface {
	OperationName() string
	Process(ctx *gin.Context) (interface{}, error)
}
