package mysql

import (
	"fmt"
	"gin-web/config"
	"gin-web/config/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

// 打开数据库链接
func OpenDB() error {
	var err error
	mysqlConfig := config.AppConfig.GetMysqlConfig()
	if !mysqlConfig.CheckInfo() {
		return fmt.Errorf("some db config info missing %+v", mysqlConfig)
	}

	DB, err = gorm.Open("mysql", makeOrmAddress(mysqlConfig))
	if err != nil {
		log.Println(err)
		return err
	}

	DB.SingularTable(true)

	return nil
}

// 拼接链接
func makeOrmAddress(config *model.MysqlConfig) string {
	return config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?charset=utf8&parseTime=true"
}

// 关闭数据库
func CloseDB() {
	if DB != nil {
		defer DB.Close()
	}
}
