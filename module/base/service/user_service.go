package service

import (
	"github.com/xiaopengkuang/gin-web/common/web"
	"fmt"
)

const (
	UserServiceName = "user"
)

type UserService struct {
}

func (u *UserService) Name() string {
	return UserServiceName
}

// 用户账号登录
func (u *UserService) API_login(request web.Request) (interface{}, error) {
	// 登录参数校验

	// token 生成 // 生成规则 token 管理规则

	return request, nil
}

// 注册新用户
func (u *UserService) API_register(request web.Request) (interface{}, error) {
	// 参数校验

	// 用户存在验证

	// 新增用户

	return nil, fmt.Errorf("thr error")
}
