package util

import "os"

const (
	EnvGoPath = "GOPATH"
	GoPathSrc = "/src/"
)

// 验证路径是否存在
func PathExist(path string) error {
	_, err := os.Stat(path)
	return err
}

// 获取gopath src路径
func GetGoPathSrc() (string, error) {
	gopath := os.Getenv(EnvGoPath)
	gopath = gopath + GoPathSrc
	err := PathExist(gopath)
	return gopath, err
}
