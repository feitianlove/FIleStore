package main

import (
	"github.com/feitianlove/FIleStore/service/account/handler"
	"github.com/feitianlove/FIleStore/service/account/proto"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
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
