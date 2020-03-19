package service

import (
	"gin-web/module/oss/operation/img"
)

const (
	ImageServiceName = "image"
)

var imageService *ImageService

type ImageService struct {
	CommonService
}

func initImageService() {
	imageService = &ImageService{CommonService: CommonService{ServiceName:ImageServiceName}}
	imageService.RegisterOperation(&img.PutImg{})
}

func GetImageService() *ImageService {
	if imageService==nil{
		initImageService()
	}
	return imageService
}
