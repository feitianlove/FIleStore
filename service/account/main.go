package main

import (
	"github.com/feitianlove/FIleStore/service/account/handler"
	"github.com/feitianlove/FIleStore/service/account/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Registry(reg),
	)
	service.Init()

	err := proto.RegisterUserServiceHandler(service.Server(), new(handler.User))
	if err != nil {
		panic(err)
	}
	err = service.Run()
	if err != nil {
		panic(err)
	}

}

//运行
//go run ./service/account/main.go --registry=consul
//go mod edit -replace google.golang.org/grpc@v1.32.0=google.golang.org/grpc@v1.26.0
