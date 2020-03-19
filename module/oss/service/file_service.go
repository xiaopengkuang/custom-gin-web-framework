package service

import (
	"gin-web/module/oss/operation/file"
)

const (
	FileServiceName = "file"
)

var fileService *FileService

type FileService struct {
	CommonService
}

func (f *FileService) Name() string {
	return FileServiceName
}

func initFileService() {
	fileService = &FileService{CommonService: CommonService{ServiceName: FileServiceName}}
	fileService.RegisterOperation(&file.PutFile{})
}

func GetFileService() *FileService {

	if fileService == nil {
		initFileService()
	}

	return fileService
}
