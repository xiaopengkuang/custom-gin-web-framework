package datadocker

import (
	"github.com/xiaopengkuang/gin-web/datadocker/orm/mysql"
	"github.com/xiaopengkuang/gin-web/datadocker/orm/redis"
)

func InitDataDocker() []error {
	errs := make([]error, 0)
	// 初始化数据库
	err := mysql.OpenDB()
	if err != nil {
		errs=append(errs,err)
	}

	// 初始化redis
	err2 := redis.InitRedis()
	if err2 != nil {
		errs=append(errs,err2)
	}

	return errs
}
