package handler

import (
	"context"
	"log"
	u "micro-simple/user-srv/model/user"
	s "micro-simple/user-srv/proto/user"
)

type Service struct{}

var (
	userService u.Service
)

func Init() {
	var err error
	userService, err = u.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

func (e *Service) QueryUserByName(ctx context.Context, req *s.Request, rsp *s.Response) error {
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &s.Error{
			Code:   500,
			Detail: err.Error(),
		}
		return nil
	}

	rsp.User = user
	rsp.Success = true
	return nil
}
