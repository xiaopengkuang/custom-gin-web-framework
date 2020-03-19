package service

import (
	"fmt"
	"gin-web/common/web"
)

const (
	UserServiceName = "user"
)

type UserService struct {
}

func (u *UserService) Name() string {
	return UserServiceName
}

func (u *UserService) API_test(request web.Request) (interface{}, error) {
	return request, nil
}

func (u *UserService) API_get(request web.Request) (interface{}, error) {

	return nil, fmt.Errorf("thr error")
}
