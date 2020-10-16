package main

import (
	"fmt"
	"micro-simple/basic"
	"micro-simple/basic/config"
	"micro-simple/user-web/handler"

	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

// main
func main() {
	// 初始化配置
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 创建新服务
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name("mu.micro.book.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":9000"),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(func(c *cli.Context) {
			// 初始化handler
			handler.Init()
		}),
	); err != nil {
		log.Fatal(err)
	}

	// 注册登录接口
	service.HandleFunc("/user/login", handler.Login)
	// 注册退出接口
	service.HandleFunc("/user/logout", handler.Logout)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
