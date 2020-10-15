package main

import (
	"fmt"
	"micro-simple/basic"
	"micro-simple/basic/config"
	"micro-simple/user-service/handler"
	"micro-simple/user-service/model"

	s "micro-simple/user-service/proto/user"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

// main ...
func main() {
	basic.Init()

	micReg := etcd.NewRegistry(registryOptions)

	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	service.Init(
		micro.Action(func(c *cli.Context) error {
			model.Init()

			handler.Init()
			return nil
		}),
	)

	s.RegisterUserHandler(service.Server(), new(handler.Service))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
