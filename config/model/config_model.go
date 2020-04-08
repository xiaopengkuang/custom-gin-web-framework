package model

import "github.com/xiaopengkuang/gin-web/util"

// 配置文件结构体
type AppConfig struct {
	Mysql *MysqlConfig
	Redis *RedisConfig
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// mysql配置
type MysqlConfig struct {
	Database string
	Host     string
	Port     string
	Username string
	Password string
}

// 获取mysql Config
func (a *AppConfig) GetMysqlConfig() *MysqlConfig {
	if a == nil {
		return nil
	}

	return a.Mysql
}

// 获取redis config
func (a *AppConfig) GetRedisConfig() *RedisConfig {
	if a == nil {
		return nil
	}

	return a.Redis
}

func (m *MysqlConfig) CheckInfo() bool {
	return util.DoAllArgsHasText(m.Database, m.Host, m.Port, m.Username, m.Password)
}

func (m *RedisConfig) CheckInfo() bool {
	return util.DoAllArgsHasText(m.Host, m.Port)
}
