package upload

import (
	"github.com/feitianlove/FIleStore/service/upload/config"
	uploadProto "github.com/feitianlove/FIleStore/service/upload/proto"
	"github.com/feitianlove/FIleStore/service/upload/router"
	"github.com/feitianlove/FIleStore/service/upload/rpc"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func startRpcService() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	service := micro.NewService(
		micro.Name("go.micro.service.upload"),
		micro.Registry(reg),
	)
	service.Init()
	err := uploadProto.RegisterUplaodServiceHandler(service.Server(), new(rpc.Upload))
	if err != nil {
		panic(err)
	}
}
func startApiService() {
	route := router.Router()
	err := route.Run(config.UploadServiceHost)
	if err != nil {
		panic(err)
	}
}
func main() {
	go startApiService()
	go startRpcService()
}
