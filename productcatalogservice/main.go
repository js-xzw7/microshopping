package main

import (
	"fmt"
	"log"
	"net"

	handler "productcatalogservice/handler"
	pb "productcatalogservice/proto"

	capi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const PORT = 50015
const ADDRESS = "localhost"

func main() {
	ipaddr := fmt.Sprintf("%s:%d", ADDRESS, PORT)

	//-----------------------注册到consul上-----------------------
	consulConfig := capi.DefaultConfig()
	consulClient, err := capi.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("new consul client err: %+v", err)
	}

	//即将注册的服务信息
	reg := capi.AgentServiceRegistration{
		Tags:    []string{"productcatalogservice"},
		Name:    "productcatalogservice",
		Address: ADDRESS,
		Port:    PORT,
	}

	err = consulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatalf("consul register err: %+v", err)
	}

	//---------------grpc 服务-----------------

	listen, err := net.Listen("tcp", ipaddr)
	if err != nil {
		log.Fatalf("listen err: %+v", err)
	}
	defer listen.Close()

	log.Default().Printf("tcp server statrt: %s", ipaddr)

	grpcService := grpc.NewServer()
	pb.RegisterProductCatalogServiceServer(grpcService, new(handler.ProductCatalogService))

	if err := grpcService.Serve(listen); err != nil {
		log.Fatalf("currencyservice set listen err: %+v", err)
	}

}
