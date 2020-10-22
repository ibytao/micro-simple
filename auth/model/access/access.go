package access

import (
	"fmt"
	"micro-simple/basic/config"
	"micro-simple/plugins/jwt"
	"micro-simple/plugins/redis"
	"sync"

	r "github.com/go-redis/redis"
	log "github.com/micro/go-micro/v2/logger"
)

var (
	s   *service
	ca  *r.Client
	m   sync.RWMutex
	cfg = &jwt.Jwt{}
)

// service 服务
type service struct {
}

type Service interface {
	// MakeAccessToken 生成token
	MakeAccessToken(subject *Subject) (ret string, err error)

	// GetCachedAccessToken 获取缓存的token
	GetCachedAccessToken(subject *Subject) (ret string, err error)

	// DelUserAccessToken 清除用户token
	DelUserAccessToken(token string) (err error)
}

// GetService 获取服务类
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

	err := config.C().App("jwt", cfg)
	if err != nil {
		panic(err)
	}
	log.Infof("[initCfg] 配置，cfg：%v", cfg)

	ca = redis.Redis()

	s = &service{}
}
