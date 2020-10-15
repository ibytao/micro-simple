package user

import (
	"fmt"
	proto "micro-simple/user-service/proto/user"
	"sync"
)

var (
	s *service
	m sync.RWMutex
)

type service struct {
}

type Service interface {
	QueryUserByName(userName string) (ret *proto.User, err error)
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	s = &service{}
}
